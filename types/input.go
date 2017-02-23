/*
 * Copyright (C) 2017 Meng Shi
 */

package types

import (
	"unsafe"
)

type AbstractInput struct {
    *AbstractCycle
    *AbstractFile
     input  Input
}

type Input interface {}

func NewInput() *AbstractInput {
    return &AbstractInput{}
}

var input = String{ len("input"), "input" }
var inputContext = &AbstractContext{
    input,
    nil,
    nil,
}

var inputs = String{ len("inputs"), "inputs" }
var inputCommands = []Command{

    { inputs,
      MAIN_CONF|CONF_BLOCK,
      inputsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func inputsBlock(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != INPUT_MODULE {
            continue
        }

        context := (*AbstractContext)(unsafe.Pointer(module.Context))
        if context == nil {
            continue
        }

        if handle := context.Create; handle != nil {
            this := handle(cycle)
            if cycle.SetContext(module.Index, &this) == Error {
                return "0"
            }
        }
    }

    if configure.SetModuleType(INPUT_MODULE) == Error {
        return "0"
    }

    if configure.Parse(cycle) == Error {
        return "0"
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != INPUT_MODULE {
            continue
        }

        this := (*AbstractContext)(unsafe.Pointer(module.Context))
        if this == nil {
            continue
        }

        context := cycle.GetContext(module.Index)
        if context == nil {
            continue
        }

        if init := this.Init; init != nil {
            if init(cycle, context) == "-1" {
                return "0"
            }
        }
    }

    return "0"
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
