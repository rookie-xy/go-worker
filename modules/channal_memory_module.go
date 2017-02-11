/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
	  "unsafe"
	. "worker/types"
	"fmt"
)

type AbstractMemoryChannal struct {
	*AbstractChannal
	*AbstractCycle

	name     string
	size     int
	channal  Channal
}

func NewMemoryChannal() *AbstractMemoryChannal {
	return &AbstractMemoryChannal{}
}

var memory = String{ len("memory"), "memory" }
var channalMemoryContext = &AbstractContext{
	memory,
	channalContextCreate,
	channalContextInit,
}

func channalContextCreate(cycle *AbstractCycle) unsafe.Pointer {
	return nil
}

func channalContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
	return ""
}

var (
	name = String{ len("name"), "name" }
	size = String{ len("size"), "size" }

	memoryChannal AbstractMemoryChannal
)

var channalMemoryCommands = []Command{

	{ name,
      MAIN_CONF|CONF_1MORE,
      configureSetFlag,
      0,
      unsafe.Offsetof(memoryChannal.name),
      nil },

	{ size,
      MAIN_CONF|CONF_1MORE,
      configureSetNumber,
      0,
      unsafe.Offsetof(memoryChannal.size),
      nil },

	NilCommand,
}

func configureSetNumber(configure *AbstractConfigure, command *Command, cycle *AbstractCycle) string {
	return ""
}

var channalMemoryModule = Module{
	MODULE_V1,
	CONTEXT_V1,
	unsafe.Pointer(channalMemoryContext),
	channalMemoryCommands,
	CHANNAL_MODULE,
	channalMemoryInit,
	channalMemoryMain,
}

func channalMemoryInit(cycle *AbstractCycle) int {
	return Ok
}

func channalMemoryMain(cycle *AbstractCycle) int {

	for ;; {
		fmt.Println("aaaaaaaaaaa")
	}

	return Ok
}

func init() {
	Modules = append(Modules, &channalMemoryModule)
}
