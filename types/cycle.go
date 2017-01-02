/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractCycle struct {
    option     *AbstractOption
    configure  *AbstractConfigure
    log        *AbstractLog
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

func (c *AbstractCycle) GetSetLog() *AbstractLog {
    return c.log
}
