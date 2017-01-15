/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"
import "fmt"

type AbstractCycle struct {
    option     *AbstractOption
    configure  *AbstractConfigure
    log        *AbstractLog
    context    []*unsafe.Pointer
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

    c.option = option

    return Ok
}

func (c *AbstractCycle) GetOption() *AbstractOption {
    return c.option
}

func (c *AbstractCycle) SetConfigure(configure *AbstractConfigure) int {
    if configure == nil {
        return Error
    }

    c.configure = configure

    return Ok
}

func (c *AbstractCycle) GetConfigure() *AbstractConfigure {
    return c.configure
}

func (c *AbstractCycle) SetLog(log *AbstractLog) int {
    if log == nil {
        return Error
    }

    c.log = log

    return Ok
}

func (c *AbstractCycle) GetLog() *AbstractLog {
    return c.log
}

func (c *AbstractCycle) SetContext(index uint, pointer *unsafe.Pointer) int {
    if index < 0 || pointer == nil {
        return Error
    }

    // TODO BUG fix
    c.context = append(c.context, pointer)
    //c.context[index] = pointer
    fmt.Println("eeeeeeeeeeeeeeeeeeeee")
    return Ok
}

func (c *AbstractCycle) GetContext(index uint) *unsafe.Pointer {
    if index < 0 {
        return nil
    }

    return c.context[index]
}
