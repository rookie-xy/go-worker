/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "worker/types"
)

const (
    MEMORY_MODULE = 0x0003
)

var memoryModule = String{ len("memory_module"), "memory_module" }
var channalMemoryContext = &AbstractContext{
    memoryModule,
    nil,
    nil,
}

var	memory = String{ len("memory"), "memory" }
var channalMemoryCommands = []Command{

    { memory,
      MAIN_CONF|CONF_1MORE,
      memoryBlock,
      0,
      0,
      nil },

    NilCommand,
}

func memoryBlock(configure *AbstractConfigure, command *Command, cycle *AbstractCycle) string {
    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != MEMORY_MODULE {
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

    if configure.SetModuleType(MEMORY_MODULE) == Error {
        return "0"
    }

    if configure.Parse(cycle) == Error {
        return "0"
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != MEMORY_MODULE {
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

var channalMemoryModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(channalMemoryContext),
    channalMemoryCommands,
    CHANNAL_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &channalMemoryModule)
}
