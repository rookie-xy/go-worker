/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
	"unsafe"
	"fmt"
	. "worker/types"
)

type AbstractLocationHttp struct {
    *AbstractHttpInput
	document string
}

func NewLocationHttp() *AbstractLocationHttp {
	return &AbstractLocationHttp{}
}

var httpLocationContext = &AbstractContext{
	location,
	httpLocationContextCreate,
	httpLocationContextInit,
}

func httpLocationContextCreate(cycle *AbstractCycle) unsafe.Pointer {
	return nil
}

func httpLocationContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
	return "0"
}

var (
	document = String{ len("document"), "document" }
	httpLocation AbstractLocationHttp
)

var httpLocationCommands = []Command{

	{ document,
      MAIN_CONF|CONF_1MORE,
	  httpLocationSet,
	  0,
	  unsafe.Offsetof(httpLocation.document),
	  nil },

	NilCommand,
}

func httpLocationSet(configure *AbstractConfigure, command *Command, cycle *AbstractCycle) string {
	return ""
}

var httpLocationModule = Module{
	MODULE_V1,
	CONTEXT_V1,
	unsafe.Pointer(httpLocationContext),
	httpLocationCommands,
	HTTP_MODULE,
	httpLocationInit,
	httpLocationMain,
}

func httpLocationInit(cycle *AbstractCycle) int {
	return Ok
}

func httpLocationMain(cycle *AbstractCycle) int {

	for ;; {
		fmt.Println("aaaaaaaaaaa")
	}

	return Ok
}

func init() {
	Modules = append(Modules, &httpLocationModule)
}