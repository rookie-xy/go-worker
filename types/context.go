/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
)

type ContextCreateFunc func(cycle *Cycle) unsafe.Pointer
type ContextInitFunc func(cycle *Cycle, configure *unsafe.Pointer) string

type Context struct {
    Name    String
    Create  ContextCreateFunc
    Init    ContextInitFunc
}

func Block(cycle *Cycle, modules []*Module, modType int64, cfgType int) int {
    if cycle == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]
        if module.Type != modType {
            continue
        }

        module.CtxIndex++
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]
        if module.Type != modType {
            continue
        }

        context := (*Context)(unsafe.Pointer(module.Context))
        if context == nil {
            continue
        }

        if handle := context.Create; handle != nil {
            this := handle(cycle)
            if cycle.SetContext(module.Index, &this) == Error {
                return Error
            }
        }
    }

    configure := cycle.GetConfigure()
    if configure == nil {
        return Error
    }

    if configure.SetModuleType(modType) == Error {
        return Error
    }

    if configure.SetCommandType(cfgType) == Error {
        return Error
    }

    if configure.Materialized(cycle, modules) == Error {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]
        if module.Type != modType {
            continue
        }

        this := (*Context)(unsafe.Pointer(module.Context))
        if this == nil {
            continue
        }

        context := cycle.GetContext(module.Index)
        if context == nil {
            continue
        }

        if init := this.Init; init != nil {
            if init(cycle, context) == "-1" {
                return Error
            }
        }
    }

    return Ok
}
