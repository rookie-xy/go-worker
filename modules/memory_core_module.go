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
    return nil
}

func coreContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    return ""
}

var (
    name = String{ len("name"), "name" }
    size = String{ len("size"), "size" }
    coreMemory AbstractMemoryCore
)

var coreMemoryCommands = []Command{

    { name,
      MAIN_CONF|CONF_1MORE,
      configureSetFlag,
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

func configureSetNumber(configure *AbstractConfigure, command *Command, cycle *AbstractCycle) string {
    value := configure.GetValue()
    fmt.Println(value)
    return ""
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