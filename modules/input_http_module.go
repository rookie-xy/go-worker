/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
	"unsafe"
	"fmt"
	. "worker/types"
)

const (
    HTTP_MODULE = 0x0001
)

type AbstractHttpInput struct {
	*AbstractCycle
	*AbstractFile

	listen  string
	location map[string]interface{}
	input    Input
}

func NewHttpInput() *AbstractHttpInput {
	return &AbstractHttpInput{}
}

var http = String{ len("http"), "http" }
var inputHttpContext = &AbstractContext{
	http,
	inputHttpContextCreate,
	inputHttpContextInit,
}

func inputHttpContextCreate(cycle *AbstractCycle) unsafe.Pointer {
	return nil
}

func inputHttpContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
	return "0"
}

var (
	listen = String{ len("listen"), "listen" }
	location = String{ len("location"), "location" }
	inputHttp AbstractHttpInput
)

var inputHttpCommands = []Command{

	{ listen,
      MAIN_CONF|CONF_1MORE,
	  configureSetFlag,
	  0,
	  unsafe.Offsetof(inputHttp.listen),
	  nil },

	{ location,
      MAIN_CONF|CONF_1MORE,
	  locationBlock,
	  0,
	  unsafe.Offsetof(inputHttp.location),
	  nil },

	NilCommand,
}

func locationBlock(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
	return ""
}

var inputHttpModule = Module{
	MODULE_V1,
	CONTEXT_V1,
	unsafe.Pointer(inputHttpContext),
	inputHttpCommands,
	INPUT_MODULE,
	inputHttpInit,
	inputHttpMain,
}

func inputHttpInit(cycle *AbstractCycle) int {
	return Ok
}

func inputHttpMain(cycle *AbstractCycle) int {

	for ;; {
		fmt.Println("aaaaaaaaaaa")
	}

	return Ok
}

func init() {
	Modules = append(Modules, &inputHttpModule)
}