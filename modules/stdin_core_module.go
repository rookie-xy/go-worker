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

func coreStdinContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
/*
    log := cycle.GetLog()

    core := (*AbstractStdinCore)(unsafe.Pointer(configure))
    if core.status == true {
        core.AbstractFile = NewFile(log)
    }
    */

    return "0"
}

var (
    status = String{ len("status"), "status" }
    channal = String{ len("channal"), "channal" }
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
    p := config
fmt.Println(p)
    fp := (*bool)(unsafe.Pointer(uintptr(*p) + command.Offset))
    fmt.Println("meililiiiiiiiiiiiiiiiiiiii", *fp)
    fmt.Println(command.Offset)

    flag := configure.GetValue()
    if flag == true {
        fmt.Println("mengshiiiiiiiiiiiiiiiii")
        *fp = true
    } else if flag == false {
        fmt.Println("zhangyueeeeeeeeeeeeeeeeeeee")
        *fp = false
    } else {
        fmt.Println("eeeeeeeeeeeeeeeeeeeeeeee")
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
    p := config
    fp := (*string)(unsafe.Pointer(uintptr(*p) + command.Offset))
    fmt.Printf("stringggggggggggggggggggggg: %s\n", *fp)
    *fp = "hahahahha"
    fmt.Println(command.Offset)

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
    fmt.Println("yyyyyyyyyyyyyyyyyyyyyyfdsafdsf")
    c := cycle.GetContext(10)
k := (*AbstractStdinCore)(unsafe.Pointer(uintptr(*c)))
    fmt.Println(k.channal)
    return Ok
}

func coreStdinMain(cycle *AbstractCycle) int {
    return Ok
}

func init() {
    Modules = append(Modules, &coreStdinModule)
}