/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
	"unsafe"
	"fmt"
	. "worker/types"
)

type AbstractStdoutInput struct {
	*AbstractCycle
	*AbstractFile

	status   bool
	channal  string
	output   Output
}

func NewStdoutInput() *AbstractStdoutInput {
	return &AbstractStdoutInput{}
}

var stdout = String{ len("stdout"), "stdout" }
var outputStdoutContext = &AbstractContext{
	stdout,
	outputStdoutContextCreate,
	outputStdoutContextInit,
}

func outputStdoutContextCreate(cycle *AbstractCycle) unsafe.Pointer {
	return nil
}

func outputStdoutContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    log := cycle.GetLog()

	output := (*AbstractStdoutInput)(unsafe.Pointer(configure))
	if output.status == true {
		output.AbstractFile = NewFile(log)
	}

	return "0"
}

var (
	outputStatus = String{ len("status"), "status" }
	outputChannal = String{ len("channal"), "channal" }
	outputStdout AbstractStdoutInput
)

var outputStdoutCommands = []Command{

	{ outputStatus,
      MAIN_CONF|CONF_1MORE,
      configureSetFlag,
      0,
      unsafe.Offsetof(outputStdout.status),
      nil },

	{ outputChannal,
      MAIN_CONF|CONF_1MORE,
      configureSetFlag,
      0,
      unsafe.Offsetof(outputStdout.channal),
      nil },

	NilCommand,
}

var outputStdoutModule = Module{
	MODULE_V1,
	CONTEXT_V1,
	unsafe.Pointer(outputStdoutContext),
	outputStdoutCommands,
	OUTPUT_MODULE,
	outputStdoutInit,
	outputStdoutMain,
}

func outputStdoutInit(cycle *AbstractCycle) int {
	return Ok
}

func outputStdoutMain(cycle *AbstractCycle) int {

	for ;; {
		fmt.Println("aaaaaaaaaaa")
	}

	return Ok
}

func init() {
	Modules = append(Modules, &outputStdoutModule)
}