/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "fmt"
)

type AbstractConfigure struct {
    *AbstractLog
    *AbstractFile
     resource   string
     fileName   string
     configure  Configure
}

type Configure interface {
    Parse() int
    ReadToken() int
}

func NewConfigure(log *AbstractLog) *AbstractConfigure {
    return &AbstractConfigure{
        AbstractLog  : log,
        AbstractFile : NewFile(log),
    }
}

func (c *AbstractConfigure) SetName(file string) int {
    if file == "" {
        return Error
    }

    if c.AbstractFile.SetName(file) == Error {
        return Error
    }

    return Ok
}

func (c *AbstractConfigure) GetName() string {
    return c.AbstractFile.GetName()
}

func (c *AbstractConfigure) SetFileName(fileName string) int {
    if fileName == "" {
        return Error
    }

    c.fileName = fileName

    return Ok
}

func (c *AbstractConfigure) GetFileName() string {
    return c.fileName
}

func (c *AbstractConfigure) SetFileType(fileType string) int {
    if fileType == "" {
        return Error
    }

    if c.SetName(fileType) == Error {
        return Error
    }

    return Ok
}

func (c *AbstractConfigure) GetFileType() string {
    if fileType := c.GetName(); fileType != "" {
        return fileType
    }

    return ""
}

func (c *AbstractConfigure) SetFile(action Action) int {
    if action == nil {
        return Error
    }

    if c.AbstractFile.Set(action) == Error {
        return Error
    }

    return Ok
}

func (c *AbstractConfigure) GetFile() Action {
    if file := c.AbstractFile.Get(); file != nil {
        return file
    }

    return nil
}

func (c *AbstractConfigure) SetResource(resource string) int {
    if resource == "" {
        return Error
    }

    c.resource = resource

    return Ok
}

func (c *AbstractConfigure) GetResource(resource string) string {
    return c.resource
}

func (c *AbstractConfigure) Get() Configure {
    log := c.AbstractLog.Get()

    file := c.AbstractFile.Get()
    if file == nil {
        file = NewFile(c.AbstractLog)
    }

    if file.Open(c.resource) == Error {
        log.Error("configure open file error")
        return nil
    }

    if file.Read() == Error {
        log.Error("configure read file error")
        return nil
    }

    if content := file.Type().GetContent(); content != nil {
        c.content = content
    } else {
        log.Warn("not found content: %d\n", 10)
    }

    return c.configure
}

func (c *AbstractConfigure) Set(configre Configure) int {
    if configre == nil {
        return Error
    }

    c.configure = configre

    return Ok
}

func (c *AbstractConfigure) Parse() int {
    fmt.Println("configure parse")
    return Ok
}

func (c *AbstractConfigure) ReadToken() int {
    fmt.Println("configure read token")
    return Ok
}
