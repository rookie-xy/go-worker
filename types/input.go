/*
 * Copyright (C) 2017 Meng Shi
 */

package types

import "unsafe"

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

func inputsBlock(cf *AbstractConfigure, command *Command, conf interface{}) string {
    return ""
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
