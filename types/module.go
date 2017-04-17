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

type Module struct {
    CtxIndex   uint
    Index      uint
    Context    unsafe.Pointer
    Commands   []Command
    Type       int64
    Init       InitFunc
    Main       MainFunc
}

var Modules []*Module

func Load(modules []*Module, module *Module) []*Module {
    if modules == nil && module == nil {
        return nil
    }

    modules = append(modules, module)

    return modules
}

func Start(modules []*Module) int {
    fmt.Println("start")
    return Ok
}

func Stop(modules []*Module) int {
    return Ok
}

func Reload(modules []*Module) int {
    if Stop(modules) == Error {
        return Error
    }

    Start(modules)

    fmt.Println("reload")

    return Ok
}

func Manager(modules []*Module) int {
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

    go f(cycle)

    return Ok
}

func (f MainFunc) Stop(c *Cycle, e *Event) int {
    c.Event = e
    return Ok
}
