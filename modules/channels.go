/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

var channal = String{ len("channal"), "channal" }
var channalContext = &Context{
    channal,
    nil,
    nil,
}

var channals = String{ len("channals"), "channals" }
var channalCommands = []Command{

    { channals,
      MAIN_CONFIG|CONFIG_BLOCK,
      channalsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func channalsBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    if cycle == nil {
        return Error
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != CHANNAL_MODULE {
            continue
        }

        module.CtxIndex++
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != CHANNAL_MODULE {
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

    if configure.SetModuleType(CHANNAL_MODULE) == Error {
        return Error
    }

    if configure.SetCommandType(CHANNAL_CONFIG) == Error {
        return Error
    }

    if configure.Materialized(cycle) == Error {
        return Error
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != CHANNAL_MODULE {
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

var channalModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(channalContext),
    channalCommands,
    CONFIG_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &channalModule)
}