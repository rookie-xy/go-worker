package modules

import (
	"unsafe"
	"fmt"
)

/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
    "unsafe"
    "fmt"
)

type AbstractInput struct {
    *AbstractCycle
    *AbstractFile
     stdin  bool
     input  Input
}

type Input interface {}

func NewInput() *AbstractInput {
    return &AbstractInput{}
}

type inputContext struct {
    *AbstractContext
}

var user = String{ len("user"), "user" }
var inputStdinContext = &AbstractContext{
    user,
    inputContextCreate,
    inputContextInit,
}

func inputContextCreate(cycle *AbstractCycle) unsafe.Pointer {
    return nil
}

func inputContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    return ""
}

func (i *inputContext) Init(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    log := cycle.GetLog()

    input := (*AbstractInput)(unsafe.Pointer(configure))
    if input.stdin == true {
        input.AbstractFile = NewFile(log)
    }

    return "0"
}

var (
    stdin = String{ len("stdin"), "stdin" }
    input AbstractInput
)

var inputCommands = []Command{

	{ stdin,
		0,
		configureSetFlag,
		0,
		unsafe.Offsetof(input.stdin),
		nil },

	NilCommand,
}

func configureSetFlag(cf *AbstractConfigure, command *Command, conf interface{}) string {
	return ""
}

var inputModule = Module{
	MODULE_V1,
	CONTEXT_V1,
	unsafe.Pointer(inputStdinContext),
	inputCommands,
	CONFIG_MODULE,
	inputInit,
	inputMain,
}

func inputInit(cycle *AbstractCycle) int {
	return Ok
}

func inputMain(cycle *AbstractCycle) int {

	for ;; {
		fmt.Println("aaaaaaaaaaa")
	}

	return Ok
}

func init() {
	Modules = append(Modules, &inputModule)
}