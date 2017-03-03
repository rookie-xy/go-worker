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
    coreStatus = String{ len("state"), "state" }
    coreChannal = String{ len("pull"), "pull" }
    coreStdout AbstractStdoutCore
)

var coreStdoutCommands = []Command{

    { coreStatus,
      STDOUT_CONFIG|CONFIG_FLAG,
      configureSetFlag,
      0,
      unsafe.Offsetof(coreStdout.status),
      nil },

    { coreChannal,
      STDOUT_CONFIG|CONFIG_FLAG,
      configureSetString,
      0,
      unsafe.Offsetof(coreStdout.channal),
      nil },

    NilCommand,
}

func configureSetFlags(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
    if config == nil {
        return "0"
    }

    pointer := (*bool)(unsafe.Pointer(uintptr(*config) + command.Offset))
    if pointer == nil {
        return "0"
    }

    flag := configure.GetValue()
    if flag == true {
        *pointer = true
    } else if flag == false {
        *pointer = false
    } else {
        return "-1"
    }

    /*
    if command.Post != nil {
        post := command.Post.(DvrConfPostType);
        post.Handler(cf, post, *p);
    }
    */

    return ""
}

func configureSetStrings(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
    pointer := (*string)(unsafe.Pointer(uintptr(*config) + command.Offset))
    if pointer == nil {
        return "0"
    }

    strings := configure.GetValue()
    if strings == nil {
        return "0"
    }

    fmt.Printf("configureSetStringsssssssssss: %s\n", *pointer)
    *pointer = strings.(string)

    return "0"
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
