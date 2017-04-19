/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
    "fmt"
)

type InitFunc func(cycle *Cycle) int
type MainFunc func(cycle *Cycle) int
type ExitFunc func(cycle *Cycle) int

type Module struct {
    CtxIndex   uint
    Index      uint
    Context    unsafe.Pointer
    Commands   []Command
    Type       int64
    Init       InitFunc
    Main       MainFunc
    Exit       ExitFunc
}

var Modules []*Module

func Load(modules []*Module, module *Module) []*Module {
    if modules == nil && module == nil {
        return nil
    }

    modules = append(modules, module)

    return modules
}

func Init(m []*Module, c *Cycle) int {
    for i := 0; m[i] != nil; i++ {
        module := m[i]

        if module.Init != nil {
            if module.Init(c) == Error {
                return Error
            }
        }
    }

    return Ok
}

func Main(m []*Module, c *Cycle) int {
    for i := 0; m[i] != nil; i++ {
        module := m[i]

        if main := module.Main; main != nil {
	           if main.Start(c) == Error {
                return Error
            }
        }
    }

    return Ok
}

func Exit(m []*Module, c *Cycle) int {
    for i := 0; m[i] != nil; i++ {
        module := m[i]

        if module.Exit != nil {
            if module.Exit(c) == Error {
                return Error
            }
        }
    }

    return Ok
}

func StartConfigModules(m []*Module, c *Cycle) int {
    modules := GetSomeModules(m, CONFIG_MODULE)
    if modules == nil {
        return Error
    }

    Main(modules, c)

    return Ok
}

func StopConfigModules(m []*Module, c *Cycle) int {
    modules := GetSomeModules(m, CONFIG_MODULE)
    if modules == nil {
        return Error
    }

    Exit(modules, c)
    // TODO clear context and other data

    return Ok
}

func ReloadModules(m []*Module, c *Cycle, flag int64) int {
    modules := GetSomeModules(m, flag)
    if modules == nil {
        return Error
    }

    Exit(modules, c)

    if Init(modules, c) == Error {
        return Error
    }

    Main(modules, c)

    fmt.Println("reload")

    return Ok
}

func Manager(m []*Module, c *Cycle) int {
    return Ok
}

func GetSomeModules(mod []*Module, modType int64) []*Module {
    var modules []*Module

    for m := 0; mod[m] != nil; m++ {
        module := mod[m]

        if module.Type == modType {
            modules = Load(modules, module)
        }
    }

    modules = Load(modules, nil)

    return modules
}

func GetSpacModules(mod []*Module) []*Module {
    var modules []*Module

    for m := 0; mod[m] != nil; m++ {
        module := mod[m]

        if module.Type == SYSTEM_MODULE ||
           module.Type == CONFIG_MODULE {
            continue
        }

        modules = Load(modules, module)
    }

    modules = Load(modules, nil)

    return modules
}

func GetPartModules(mod []*Module, modType int64) []*Module {
    if mod == nil || len(mod) <= 0 {
        return nil
    }

    switch modType {

    case SYSTEM_MODULE:
        modules := GetSomeModules(mod, modType)
        if modules != nil {
            return modules
        }

    case CONFIG_MODULE:
        modules := GetSomeModules(mod, modType)
        if modules != nil {
            return modules
        }
    }

    var modules []*Module

    modType = modType >> 28

    for m := 0; mod[m] != nil; m++ {
        module := mod[m]
        moduleType := module.Type >> 28

        if moduleType == modType {
            modules = Load(modules, module)
        }
    }

    modules = Load(modules, nil)

    return modules
}

func (f MainFunc) Start(cycle *Cycle) int {
    if f == nil {
        return Error
    }

//    cycle.Routine.Go()

    go f(cycle)

    return Ok
}
