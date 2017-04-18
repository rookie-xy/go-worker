/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
    "os"
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
     //modules  []*Module

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
/*
func (c *Cycle) SetModules(modules []*Module) int {
    if modules == nil || len(modules) <= 0 {
        return Error
    }

    c.modules = modules

    return Ok
}

func (c *Cycle) GetModules() []*Module {
    if c.modules == nil || len(c.modules) <= 0 {
        return nil
    }

    return c.modules
}

func (c *Cycle) GetSpacModules() []*Module {
    return GetSpacModules(c.modules)
}

func (c *Cycle) GetSomeModules(moduleType int64) []*Module {
    return GetSomeModules(c.modules, moduleType)
}

func (c *Cycle) GetPartModules(moduleType int64) []*Module {
    return GetPartModules(c.modules, moduleType)
}
*/
func (c *Cycle) Start(modules []*Module) int {
    if cycle := c.cycle; cycle != nil {
        if cycle.Start(modules) == Error {
            return Error
        }

        return Ok
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]

        if main := module.Main; main != nil {
            if main.Start(c) == Error {
                return Error
            }
        }
    }

    return Ok
}
/*
func (c *Cycle) Stop() int {
    if cycle := c.cycle; cycle != nil {
        if cycle.Stop() == Error {
            return Error
        }

        return Ok
    }

    //kill := NewEvent()

    for m := 0; c.modules[m] != nil; m++ {
        module := c.modules[m]

        if main := module.Main; main != nil {
            //fmt.Printf("WANGWANGSSSSSSSSSSSSSSSSSSSSSSSS: %X\n", module.Type)
            //kill.SetOpcode(int(module.Type))
            //main.Stop(c, kill)
        }
    }

    return Ok
}

func (c *Cycle) Reload() int {
    c.Stop()
fmt.Println("stop")
    c.System()
    fmt.Println("system")

    c.ConfigureInit()
    fmt.Println("configure init")

    if c.Routine == nil {
        c.Routine = NewRoutine()
    }

    modules := c.GetSomeModules(CONFIG_MODULE)
    if modules == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]

        if init := module.Init; init != nil {
            if init(c) == Error {
                return Error
            }
        }
    }

    if c.Start(modules) == Error {
        return Error
    }

    return Ok
}
*/

func (c *Cycle) Init(m []*Module) int {
    modules:= GetSomeModules(m, SYSTEM_MODULE)
    if modules == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]

        if module.Init != nil {
            if module.Init(c) == Error {
                os.Exit(1)
            }
        }

        if main := module.Main; main != nil {
	           if main.Start(c) == Error {
                os.Exit(2)
            }
        }
    }

    //configure := c.Configure
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

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]

        if init := module.Init; init != nil {
            if init(c) == Error {
                return Error
            }
        }
    }

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
                if Start(m, c) == Error {
                    return Error
                }

            case STOP:
                if Stop(m, c) == Error {
                    return Error
                }

            case RELOAD:
                if Reload(m, c) == Error {
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
}

func (c *Cycle) Exit() {
    return
}
/*
func (c *Cycle) System() int {
    modules:= c.GetSomeModules(SYSTEM_MODULE)
    if modules == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]

        if module.Init != nil {
            if module.Init(c) == Error {
                os.Exit(1)
            }
        }

        if main := module.Main; main != nil {
	           if main.Start(c) == Error {
                os.Exit(2)
            }
        }
    }

    return Ok
}
*/
/*
func (c *Cycle) ConfigureInit() int {
    select {

    case e := <- c.Event:
        if op := e.GetOpcode(); op != LOAD {
            return Ignore
        }
    }

    if c.Block(c, CONFIG_MODULE, CONFIG_BLOCK) == Error {
        return Error
    }

    return Ok
}
*/
