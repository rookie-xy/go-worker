/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    INPUT_MODULE = 0x30000000
    INPUT_CONFIG = 0x00300000
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
    cycle.Configure.Block(INPUT_MODULE, INPUT_CONFIG)
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
    Modules = Load(Modules, &inputModule)
}