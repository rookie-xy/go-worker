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
    fmt.Println("reload")
    return Ok
}

func Manager(modules []*Module) int {
    return Ok
}

func (f MainFunc) Start(cycle *Cycle) int {
    if f == nil {
        return Error
    }

    go f(cycle)
    //cycle.Routine.Go(0, f, cycle)

    return Ok
}

func (f MainFunc) Stop() int {
    return Ok
}
