/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    FILTER_MODULE = 0x696600000000
    FILTER_CONFIG = 0x10000000
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
      MAIN_CONFIG|CONFIG_BLOCK,
      filterBlock,
      0,
      0,
      nil },

    NilCommand,
}

func filterBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    cycle.Configure.Block(FILTER_MODULE, FILTER_CONFIG)
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
}

func init() {
    //Modules = append(Modules, &filterModule)
}