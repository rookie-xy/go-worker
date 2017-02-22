/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "worker/types"
)

type AbstractStdoutCore struct {
    *AbstractCycle
    *AbstractFile

     status   bool
     channal  string
}

func NewStdoutCore() *AbstractStdoutCore {
    return &AbstractStdoutCore{}
}

var stdoutCore = String{ len("stdout_core"), "stdout_core" }
var coreStdoutContext = &AbstractContext{
    stdoutCore,
    coreStdoutContextCreate,
    coreStdoutContextInit,
}

func coreStdoutContextCreate(cycle *AbstractCycle) unsafe.Pointer {
    return nil
}

func coreStdoutContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    log := cycle.GetLog()

    core := (*AbstractStdoutCore)(unsafe.Pointer(configure))
    if core.status == true {
        core.AbstractFile = NewFile(log)
    }

    return "0"
}

var (
    coreStatus = String{ len("status"), "status" }
    coreChannal = String{ len("channal"), "channal" }
    coreStdout AbstractStdoutCore
)

var coreStdoutCommands = []Command{

    { coreStatus,
      MAIN_CONF|CONF_1MORE,
      configureSetFlag,
      0,
      unsafe.Offsetof(coreStdout.status),
      nil },

    { coreChannal,
      MAIN_CONF|CONF_1MORE,
      configureSetFlag,
      0,
      unsafe.Offsetof(coreStdout.channal),
      nil },

    NilCommand,
}

var coreStdoutModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(coreStdoutContext),
    coreStdoutCommands,
    STDOUT_MODULE,
    coreStdoutInit,
    coreStdoutMain,
}

func coreStdoutInit(cycle *AbstractCycle) int {
    return Ok
}

func coreStdoutMain(cycle *AbstractCycle) int {
    return Ok
}

func init() {
    Modules = append(Modules, &coreStdoutModule)
}