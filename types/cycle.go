/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
    "sync"
)

type Cycle struct {
    *Log
    *Routine
    *Option
    *Configure

    *Input
    *Channel
    *Output
    *Codec
    *Filter

     sync.Mutex

     context  [1024]*unsafe.Pointer
     cycle    CycleIf
}

type Handle interface {
    Get() int
    Set() int

    GetType() unsafe.Pointer
}

type CycleIf interface {
    Start(module []*Module) int
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

func (c *Cycle) Start(modules []*Module) int {
    if cycle := c.cycle; cycle != nil {
        if cycle.Start(modules) == Error {
            return Error
        }

        return Ok
    }

    Main(modules, c)

    return Ok
}

func (c *Cycle) Init(m []*Module) int {
    modules:= GetSomeModules(m, SYSTEM_MODULE)
    if modules == nil {
        return Error
    }

    Init(modules, c)

    Main(modules, c)

    if c.Configure == nil {
        c.Configure = NewConfigure(c)
    }

    select {

    case e := <- c.Event:
        if op := e.GetOpcode(); op != LOAD {
            return Ignore
        }
    }

    if Block(c, m, CONFIG_MODULE, CONFIG_BLOCK) == Error {
        return Error
    }

    return Ok
}

func (c *Cycle) Main(m []*Module) int {
    if c.Routine == nil {
        c.Routine = NewRoutine()
    }

    modules := GetSpacModules(m)
    if modules == nil && c == nil {
        return Error
    }

    Init(modules, c)

    if c.Start(modules) == Error {
        return Error
    }

    return Ok
}

func (c *Cycle) Monitor(m []*Module) int {
    for {
        select {

        case event := <- c.Event:
            opcode := event.GetOpcode()

            switch opcode {

            case START:
                if StartConfigModules(m, c) == Error {
                    return Error
                }

            case STOP:
                if StopConfigModules(m, c) == Error {
                    return Error
                }

            case RELOAD:
                if ReloadModules(m, c, CONFIG_MODULE) == Error {
                    return Error
                }
            }
        }
    }
/*
    if routine := c.Routine; routine != nil {
        if routine.Monitor() == Error {
            return Error
        }
    }
    */

    return Ok
}

func (c *Cycle) Exit(m []*Module) {
    Exit(m, c)
    return
}
