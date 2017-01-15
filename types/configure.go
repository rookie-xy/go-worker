/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "fmt"

type AbstractConfigure struct {
    *AbstractFile
    fileName   string
    configure  Configure
}

type Configure interface {
    Parse() int
    ReadToken() int
}

func NewConfigure() *AbstractConfigure {
    return &AbstractConfigure{
        AbstractFile : NewFile(),
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

func (c *AbstractConfigure) SetFile(file File) int {
    if file == nil {
        return Error
    }

    if c.AbstractFile.SetFile(file) == Error {
        return Error
    }

    return Ok
}

func (c *AbstractConfigure) GetFile() File {
    if file := c.AbstractFile.GetFile(); file != nil {
        return file
    }

    return nil
}

func (c *AbstractConfigure) Get() Configure {
    file := c.AbstractFile.GetFile()
    if file == nil {
        file = NewFile()
    }

    if file.Open() == Error {
        return nil
    }

    if file.Read() == Error {
        return nil
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
