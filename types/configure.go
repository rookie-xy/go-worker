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
    Parse() int
    ReadToken() int
}

func NewConfigure() *AbstractConfigure {
    return &AbstractConfigure{}
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

func (c *AbstractConfigure) Get() Configure {
    file := c.AbstractFile
    if file == nil {
        return nil
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

