/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

type AbstractCycle struct {
    *AbstractOption
    *AbstractConfigure
    *AbstractLog
    context  [1024]*unsafe.Pointer
}

type Cycle interface {
    Get() int
}

func NewCycle() *AbstractCycle {
    return &AbstractCycle{}
}

func (c *AbstractCycle) SetOption(option *AbstractOption) int {
    if option == nil {
        return Error
    }

    c.AbstractOption = option

    return Ok
}

func (c *AbstractCycle) GetOption() *AbstractOption {
    return c.AbstractOption
}

func (c *AbstractCycle) SetConfigure(configure *AbstractConfigure) int {
    if configure == nil {
        return Error
    }

    c.AbstractConfigure = configure

    return Ok
}

func (c *AbstractCycle) GetConfigure() *AbstractConfigure {
    return c.AbstractConfigure
}

func (c *AbstractCycle) SetLog(log *AbstractLog) int {
    if log == nil {
        return Error
    }

    c.AbstractLog = log

    return Ok
}

func (c *AbstractCycle) GetLog() *AbstractLog {
    return c.AbstractLog
}

func (c *AbstractCycle) SetContext(index uint, pointer *unsafe.Pointer) int {
    if index < 0 || pointer == nil {
        return Error
    }

    c.context[index] = pointer

    return Ok
}

func (c *AbstractCycle) GetContext(index uint) *unsafe.Pointer {
    if index < 0 {
        return nil
    }

    return c.context[index]
}
