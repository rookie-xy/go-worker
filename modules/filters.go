/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    FILTER_MODULE = 0x60000000
    FILTER_CONFIG = MAIN_CONFIG|CONFIG_BLOCK
)

var filter = String{ len("filter"), "filter" }
var filterContext = &Context{
    filter,
    nil,
    nil,
}

var filters = String{ len("filters"), "filters" }
var filterCommands = []Command{

    { filters,
      FILTER_CONFIG,
      filterBlock,
      0,
      0,
      nil },

    NilCommand,
}

func filterBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    if nil == cycle {
        return Error
    }

    flag := FILTER_CONFIG|CONFIG_ARRAY
    cycle.Block(cycle, FILTER_MODULE, flag)

    return Ok
}

var filterModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(filterContext),
    filterCommands,
    CONFIG_MODULE,
    nil,
    nil,
    nil,
}

func init() {
    Modules = Load(Modules, &filterModule)
}