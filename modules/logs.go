/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    LOG_MODULE = 0x50000000
    LOG_CONFIG = LOG_MODULE|0x00000001
)

var log = String{ len("log"), "log" }
var logContext = &Context{
    log,
    nil,
    nil,
}

var logs = String{ len("logs"), "logs" }
var logCommands = []Command{

    { logs,
      MAIN_CONFIG|CONFIG_BLOCK,
      logsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func logsBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    cycle.Configure.Block(LOG_MODULE, LOG_CONFIG)
    return Ok
}

var logModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(logContext),
    logCommands,
    CONFIG_MODULE,
    nil,
    nil,
}

func init() {
    //Modules = Load(Modules, &logModule)
}