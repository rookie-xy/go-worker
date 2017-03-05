/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "worker/types"
"fmt"
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
    stdoutCore := NewStdoutCore()
    if stdoutCore == nil {
        return nil
    }

    stdoutCore.status = false
    stdoutCore.channal = "zhangyue"

    return unsafe.Pointer(stdoutCore)
}

func coreStdoutContextInit(cycle *AbstractCycle, context *unsafe.Pointer) string {
    log := cycle.GetLog()
    this := (*AbstractStdoutCore)(unsafe.Pointer(uintptr(*context)))
    if this == nil {
        log.Error("coreStdoutContextInit error")
        return "0"
    }

    fmt.Println(this.channal)

    return "0"
}

var (
    coreStatus = String{ len("status"), "status" }
    coreChannal = String{ len("channal"), "channal" }
    coreStdout AbstractStdoutCore
)

var coreStdoutCommands = []Command{

    { coreStatus,
      STDOUT_CONFIG,
      SetFlag,
      0,
      unsafe.Offsetof(coreStdout.status),
      nil },

    { coreChannal,
      STDOUT_CONFIG,
      SetString,
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
