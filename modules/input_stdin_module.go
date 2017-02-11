/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
      "fmt"
	. "worker/types"
)

type AbstractStdinInput struct {
    *AbstractCycle
    *AbstractFile

	 status   bool
	 channal  string
     input    Input
}

func NewStdinInput() *AbstractStdinInput {
    return &AbstractStdinInput{}
}

type inputContext struct {
    *AbstractContext
}

var stdin = String{ len("stdin"), "stdin" }
var inputStdinContext = &AbstractContext{
    stdin,
    inputStdinContextCreate,
    inputStdinContextInit,
}

func inputStdinContextCreate(cycle *AbstractCycle) unsafe.Pointer {
    return nil
}

func inputStdinContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    return ""
}

func (i *inputContext) Init(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    log := cycle.GetLog()

    input := (*AbstractStdinInput)(unsafe.Pointer(configure))
    if input.status == true {
        input.AbstractFile = NewFile(log)
    }

    return "0"
}

var (
    status = String{ len("status"), "status" }
    channal = String{ len("channal"), "channal" }
    inputStdin AbstractStdinInput
)

var inputStdinCommands = []Command{

	{ status,
      MAIN_CONF|CONF_1MORE,
      configureSetFlag,
      0,
      unsafe.Offsetof(inputStdin.status),
      nil },

	{ channal,
      MAIN_CONF|CONF_1MORE,
      configureSetFlag,
      0,
      unsafe.Offsetof(inputStdin.channal),
      nil },

	NilCommand,
}

func configureSetFlag(configure *AbstractConfigure, command *Command, cycle *AbstractCycle) string {
	return ""
}

var inputStdinModule = Module{
	MODULE_V1,
	CONTEXT_V1,
	unsafe.Pointer(inputStdinContext),
	inputStdinCommands,
	INPUT_MODULE,
	inputStdinInit,
	inputStdinMain,
}

func inputStdinInit(cycle *AbstractCycle) int {
	return Ok
}

func inputStdinMain(cycle *AbstractCycle) int {

	for ;; {
		fmt.Println("aaaaaaaaaaa")
	}

	return Ok
}

func init() {
	Modules = append(Modules, &inputStdinModule)
}