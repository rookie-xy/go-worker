/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
      "fmt"
    . "worker/types"
)

type AbstractStdinCore struct {
    *AbstractCycle
    *AbstractFile

     status   bool
     channal  string
}

func NewStdinCore() *AbstractStdinCore {
    return &AbstractStdinCore{}
}

var stdinCore = String{ len("stdin_core"), "stdin_core" }
var coreStdinContext = &AbstractContext{
    stdinCore,
    coreStdinContextCreate,
    coreStdinContextInit,
}

func coreStdinContextCreate(cycle *AbstractCycle) unsafe.Pointer {
    stdinCore := NewStdinCore()
    if stdinCore == nil {
        return nil
    }

    stdinCore.status = false
    stdinCore.channal = "mengshi"

    return unsafe.Pointer(stdinCore)
}

func coreStdinContextInit(cycle *AbstractCycle, context *unsafe.Pointer) string {
    log := cycle.GetLog()
    this := (*AbstractStdinCore)(unsafe.Pointer(uintptr(*context)))
    if this == nil {
        log.Error("coreStdinContextInit error")
        return "0"
    }

    fmt.Println(this.channal)

    return "0"
}

var (
    status = String{ len("status"), "status" }
    channal = String{ len("channal"), "channal" }
    coreStdin AbstractStdinCore
)

var coreStdinCommands = []Command{

    { status,
      STDIN_CONFIG,
      SetFlag,
      0,
      unsafe.Offsetof(coreStdin.status),
      nil },

    { channal,
      STDIN_CONFIG,
      SetString,
      0,
      unsafe.Offsetof(coreStdin.channal),
      nil },

    NilCommand,
}

var coreStdinModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(coreStdinContext),
    coreStdinCommands,
    STDIN_MODULE,
    coreStdinInit,
    coreStdinMain,
}

func coreStdinInit(cycle *AbstractCycle) int {
    /*
    context := cycle.GetContext(10)
    if context == nil {
        return Error
    }

    this := (*AbstractStdinCore)(unsafe.Pointer(uintptr(*context)))
    if this == nil {
        return Error
    }

    fmt.Println(this.channal)
    */

    return Ok
}

func coreStdinMain(cycle *AbstractCycle) int {
    return Ok
}

func init() {
    Modules = append(Modules, &coreStdinModule)
}