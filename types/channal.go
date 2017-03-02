/*
 * Copyright (C) 2017 Meng Shi
 */

package types

import "unsafe"

type AbstractChannal struct {
    *AbstractCycle
    *AbstractFile
     channal  Channal
}

type Channal interface {}

func NewChannal() *AbstractChannal {
    return &AbstractChannal{}
}

var channal = String{ len("channal"), "channal" }
var channalContext = &AbstractContext{
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

func channalsBlock(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != CHANNAL_MODULE {
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

    if configure.SetModuleType(CHANNAL_MODULE) == Error {
        return "0"
    }

    if configure.Parse(cycle) == Error {
        return "0"
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != CHANNAL_MODULE {
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