/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

var input = String{ len("input"), "input" }
var inputContext = &Context{
    input,
    nil,
    nil,
}

var inputs = String{ len("inputs"), "inputs" }
var inputCommands = []Command{

    { inputs,
      MAIN_CONFIG|CONFIG_BLOCK,
      inputsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func inputsBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    if cycle == nil {
        return Error
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != INPUT_MODULE {
            continue
        }

        module.CtxIndex++
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != INPUT_MODULE {
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

    if configure.SetModuleType(INPUT_MODULE) == Error {
        return Error
    }

    if configure.SetCommandType(INPUT_CONFIG) == Error {
        return Error
    }

    if configure.Parse(cycle) == Error {
        return Error
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != INPUT_MODULE {
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

var inputModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(inputContext),
    inputCommands,
    CONFIG_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &inputModule)
}