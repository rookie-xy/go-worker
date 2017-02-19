/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
	. "worker/types"
)

const (
    STDIN_MODULE = 0x0002
)

var stdinModule = String{ len("stdin_module"), "stdin_module" }
var inputStdinContext = &AbstractContext{
    stdinModule,
    nil,
    nil,
}

var stdin = String{ len("stdin"), "stdin" }
var inputStdinCommands = []Command{

	{ stdin,
      MAIN_CONF|CONF_1MORE,
      stdinBlock,
      0,
      0,
      nil },

	NilCommand,
}

func stdinBlock(configure *AbstractConfigure, command *Command, cycle *AbstractCycle) string {
	for m := 0; Modules[m] != nil; m++ {
		module := Modules[m]
		if module.Type != STDIN_MODULE {
			continue
		}

		context := (*AbstractContext)(unsafe.Pointer(module.Context))
		if context == nil {
			continue
		}

		if handle := context.Create; handle != nil {
			this := handle(cycle)
			/*
			if *(*string)(unsafe.Pointer(uintptr(this))) == "-1" {
				return "0";
			}
			*/

			if cycle.SetContext(module.Index, &this) == Error {
				return "0"
			}
		}
	}

	if configure.SetModuleType(STDIN_MODULE) == Error {
		return "0"
	}

	if configure.Parse(cycle) == Error {
		return "0"
	}

	for m := 0; Modules[m] != nil; m++ {
		module := Modules[m]
		if module.Type != STDIN_MODULE {
			continue
		}

		this := (*AbstractContext)(unsafe.Pointer(module.Context))
		if this == nil {
			continue
		}

		context := cycle.GetContext(module.Index)
		if context == nil {
			continue
		}

		if init := this.Init; init != nil {
			if init(cycle, context) == "-1" {
				return "0"
			}
		}
	}

	return "0"
}

var inputStdinModule = Module{
	MODULE_V1,
	CONTEXT_V1,
	unsafe.Pointer(inputStdinContext),
	inputStdinCommands,
	INPUT_MODULE,
	nil,
	nil,
}

func init() {
	Modules = append(Modules, &inputStdinModule)
}