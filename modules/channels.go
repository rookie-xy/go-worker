/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    CHANNEL_MODULE = 0x686300000000
    CHANNEL_CONFIG = 0x00001000
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
    cycle.Configure.Block(CHANNEL_MODULE, CHANNEL_CONFIG)
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