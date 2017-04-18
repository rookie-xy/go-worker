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
    INPUT_CONFIG = MAIN_CONFIG|CONFIG_BLOCK
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
      INPUT_CONFIG,
      inputsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func inputsBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    if nil == cycle {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    cycle.Block(cycle, INPUT_MODULE, flag)

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
    nil,
}

func init() {
    Modules = Load(Modules, &inputModule)
}