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
    channal = String{ len("push"), "push" }
    coreStdin AbstractStdinCore
)

var coreStdinCommands = []Command{

    { status,
      MAIN_CONF|CONF_1MORE,
      configureSetFlag,
      0,
      unsafe.Offsetof(coreStdin.status),
      nil },

    { channal,
      MAIN_CONF|CONF_1MORE,
      configureSetString,
      0,
      unsafe.Offsetof(coreStdin.channal),
      nil },

    NilCommand,
}

func configureSetFlag(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
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

func configureSetString(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
    pointer := (*string)(unsafe.Pointer(uintptr(*config) + command.Offset))
    if pointer == nil {
        return "0"
    }

    strings := configure.GetValue()
    if strings == nil {
        return "0"
    }

    fmt.Printf("configureSetString: %s\n", *pointer)
    *pointer = strings.(string)

    return "0"
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