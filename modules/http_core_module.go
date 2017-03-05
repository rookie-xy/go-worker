/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "worker/types"
    "fmt"
)

const (
    LOCATION_CONFIG = 0x00010001
)

type AbstractHttpCore struct {
    *AbstractCycle
    *AbstractFile

     listen    string
     timeout   int
     location  *AbstractLocationHttp
     input     Input
}

func NewHttpCore() *AbstractHttpCore {
    return &AbstractHttpCore{}
}

var httpCore = String{ len("http_core"), "http_core" }
var coreHttpContext = &AbstractContext{
    httpCore,
    coreHttpContextCreate,
    coreHttpContextInit,
}

func coreHttpContextCreate(cycle *AbstractCycle) unsafe.Pointer {
    coreHttp := NewHttpCore()
    if coreHttp == nil {
        return nil
    }

    coreHttp.listen = "127.0.0.1:9800"
    coreHttp.timeout = 3
    coreHttp.location = nil

    return unsafe.Pointer(coreHttp)
}

func coreHttpContextInit(cycle *AbstractCycle, context *unsafe.Pointer) string {
    log := cycle.GetLog()
    this := (*AbstractHttpCore)(unsafe.Pointer(uintptr(*context)))
    if this == nil {
        log.Error("coreStdinContextInit error")
        return "0"
    }

    fmt.Println(this.listen)

    return "0"
}

var (
    listen   = String{ len("listen"), "listen" }
    timeout  = String{ len("timeout"), "timeout" }
    location = String{ len("location"), "location" }

    coreHttp  AbstractHttpCore
)

var coreHttpCommands = []Command{

    { listen,
      HTTP_CONFIG,
      SetString,
      0,
      unsafe.Offsetof(coreHttp.listen),
      nil },

    { timeout,
      HTTP_CONFIG,
      SetNumber,
      0,
      unsafe.Offsetof(coreHttp.timeout),
      nil },

    { location,
      HTTP_CONFIG,
      locationBlock,
      0,
      unsafe.Offsetof(coreHttp.location),
      nil },

    NilCommand,
}

func locationBlock(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != HTTP_MODULE {
            continue
        }

        module.CtxIndex++
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != HTTP_MODULE {
            continue
        }

        context := (*AbstractContext)(unsafe.Pointer(module.Context))
        if context == nil {
            continue
        }

        if handle := context.Create; handle != nil {
            this := handle(cycle)
            if cycle.SetContext(module.Index, &this) == Error {
                return "0"
            }
        }
    }

    if configure.SetModuleType(HTTP_MODULE) == Error {
        return "0"
    }

    if configure.SetCommandType(LOCATION_CONFIG) == Error {
        return "0"
    }

    if configure.Parse(cycle) == Error {
        return "0"
    }

    for m := 0; Modules[m] != nil; m++ {
        module := Modules[m]
        if module.Type != HTTP_MODULE {
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

var coreHttpModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(coreHttpContext),
    coreHttpCommands,
    HTTP_MODULE,
    coreHttpInit,
    coreHttpMain,
}

func coreHttpInit(cycle *AbstractCycle) int {
    return Ok
}

func coreHttpMain(cycle *AbstractCycle) int {
    return Ok
}

func init() {
    Modules = append(Modules, &coreHttpModule)
}
