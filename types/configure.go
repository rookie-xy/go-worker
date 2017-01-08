/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "fmt"

var ConfigureTypeName string = "configure_type_name"

type AbstractConfigure struct {
    file       string
    type_name  string
    configure  Configure
}

type Configure interface {
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

    c.file = file

    return Ok
}

func (c *AbstractConfigure) GetFile() string {
    return c.file
}

func (c *AbstractConfigure) SetTypeName(name string) int {
    if name == "" {
        return Error
    }

    c.type_name = name

    return Ok
}

func (c *AbstractConfigure) GetTypeName() string {
    return c.type_name
}

func (c *AbstractConfigure) Parse() int {
    fmt.Println("configure parse")
    return Ok
}

func (c *AbstractConfigure) ReadToken() int {
    fmt.Println("configure read token")
    return Ok
}

func (c *AbstractConfigure) Override() Configure {
    return c.configure
}

func (c *AbstractConfigure) Set(configre Configure) int {
    if configre == nil {
        return Error
    }

    c.configure = configre

    return Ok
}
