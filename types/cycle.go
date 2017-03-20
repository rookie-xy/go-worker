/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
)

type Cycle struct {
    *Log
    *Routine
    *Option
    *Configure

    *Input
    *Channal
    *Output
    *Codec
    *Filter

     context  [1024]*unsafe.Pointer
     cycle    CycleIf
}

type CycleIf interface {
    Start() int
    Stop() int
}

func NewCycle(log *Log) *Cycle {
    return &Cycle{
        Log:log,
    }
}

func (c *Cycle) SetOption(option *Option) int {
    if option == nil {
        return Error
    }

    c.Option = option

    return Ok
}

func (c *Cycle) GetOption() *Option {
    return c.Option
}

func (c *Cycle) SetConfigure(configure *Configure) int {
    if configure == nil {
        return Error
    }

    c.Configure = configure

    return Ok
}

func (c *Cycle) GetConfigure() *Configure {
    return c.Configure
}

func (c *Cycle) SetLog(log *Log) int {
    if log == nil {
        return Error
    }

    c.Log = log

    return Ok
}

func (c *Cycle) GetLog() *Log {
    return c.Log
}

func (c *Cycle) SetContext(index uint, pointer *unsafe.Pointer) int {
    if index < 0 || pointer == nil {
        return Error
    }

    c.context[index] = pointer

    return Ok
}

func (c *Cycle) GetContext(index uint) *unsafe.Pointer {
    if index < 0 {
        return nil
    }

    return c.context[index]
}

func (c *Cycle) Start() int {
    if cycle := c.cycle; cycle != nil {
        if cycle.Start() == Error {
            return Error
        }

        return Ok
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]

        if main := module.Main; main != nil {
            if main.Start(c) == Error {
                return Error
            }
        }
    }

    return Ok
}

func (c *Cycle) Stop() int {
    if cycle := c.cycle; cycle != nil {
        if cycle.Stop() == Error {
            return Error
        }

        return Ok
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]

        if main := module.Main; main != nil {
            main.Stop()
        }
    }

    return Ok
}
