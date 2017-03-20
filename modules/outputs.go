/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

var output = String{ len("output"), "output" }
var outputContext = &Context{
    output,
    nil,
    nil,
}

var outputs = String{ len("outputs"), "outputs" }
var outputCommands = []Command{

    { outputs,
      MAIN_CONFIG|CONFIG_BLOCK,
      outputsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func outputsBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    if cycle == nil {
        return Error
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != OUTPUT_MODULE {
            continue
        }

        module.CtxIndex++
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != OUTPUT_MODULE {
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

    if configure.SetModuleType(OUTPUT_MODULE) == Error {
        return Error
    }

    if configure.SetCommandType(OUTPUT_CONFIG) == Error {
        return Error
    }

    if configure.Materialized(cycle) == Error {
        return Error
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != OUTPUT_MODULE {
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

var outputModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(outputContext),
    outputCommands,
    CONFIG_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &outputModule)
}