/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "fmt"

type AbstractConfigure struct {
    *AbstractFile
    typeName       string
    configure      Configure
}

type Configure interface {
    Get() int      // get file from local or remote
    Parse() int
    ReadToken() int
}

func NewConfigure() *AbstractConfigure {
    return &AbstractConfigure{}
}

func (c *AbstractConfigure) SetFile(file string) int {
    if file == "" {
        return Error
    }

    if c.AbstractFile.SetName(file) == Error {
        return Error
    }

    return Ok
}

func (c *AbstractConfigure) GetFile() string {
    return c.AbstractFile.GetName()
}

func (c *AbstractConfigure) SetTypeName(name string) int {
    if name == "" {
        return Error
    }

    c.typeName = name

    return Ok
}

func (c *AbstractConfigure) GetTypeName() string {
    return c.typeName
}

func (c *AbstractConfigure) GetConfigure() Configure {
    return c.configure
}

func (c *AbstractConfigure) SetConfigure(configre Configure) int {
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

