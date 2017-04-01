/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    OUTPUT_MODULE = 0x40000000
    OUTPUT_CONFIG = 0x04000000
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
    cycle.Configure.Block(OUTPUT_MODULE, OUTPUT_CONFIG)
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
    Modules = Load(Modules, &outputModule)
}