/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
      "fmt"
    . "worker/types"
)

type AbstractMemoryCore struct {
    *AbstractChannal
    *AbstractCycle

     name     string
     size     int
}

func NewMemoryCore() *AbstractMemoryCore {
    return &AbstractMemoryCore{}
}

var memoryCore = String{ len("memory_core"), "memory_core" }
var coreMemoryContext = &AbstractContext{
    memoryCore,
    coreContextCreate,
    coreContextInit,
}

func coreContextCreate(cycle *AbstractCycle) unsafe.Pointer {
    memoryCore := NewMemoryCore()
    if memoryCore == nil {
        return nil
    }

    memoryCore.name = "memory test"
    memoryCore.size = 1024

    return unsafe.Pointer(memoryCore)
}

func coreContextInit(cycle *AbstractCycle, context *unsafe.Pointer) string {
    log := cycle.GetLog()
    this := (*AbstractMemoryCore)(unsafe.Pointer(uintptr(*context)))
    if this == nil {
        log.Error("coreStdinContextInit error")
        return "0"
    }

    fmt.Println(this.name)
    fmt.Println(this.size)

    return "0"
}

var (
    name = String{ len("name"), "name" }
    size = String{ len("size"), "size" }
    coreMemory AbstractMemoryCore
)

var coreMemoryCommands = []Command{

    { name,
      MAIN_CONF|CONF_1MORE,
      configureSetString,
      0,
      unsafe.Offsetof(coreMemory.name),
      nil },

    { size,
      MAIN_CONF|CONF_1MORE,
      configureSetNumber,
      0,
      unsafe.Offsetof(coreMemory.size),
      nil },

    NilCommand,
}

func configureSetNumber(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {

    pointer := (*int)(unsafe.Pointer(uintptr(*config) + command.Offset))
    if pointer == nil {
        return "0"
    }

    number := configure.GetValue()
    if number == nil {
        return "0"
    }

    fmt.Printf("configureSetNumber: %d\n", *pointer)
    *pointer = number.(int)

    return "0"
}

var coreMemoryModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(coreMemoryContext),
    coreMemoryCommands,
    MEMORY_MODULE,
    coreMemoryInit,
    coreMemoryMain,
}

func coreMemoryInit(cycle *AbstractCycle) int {
    return Ok
}

func coreMemoryMain(cycle *AbstractCycle) int {
    return Ok
}

func init() {
    Modules = append(Modules, &coreMemoryModule)
}