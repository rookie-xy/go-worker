/*
 * Copyright (C) 2017 Meng Shi
 */

package types

import (
    "unsafe"
)

type AbstractChannal struct {
    *AbstractCycle

     name    string
     channal  Channal
}

type Channal interface {}

func NewChannal() *AbstractChannal {
    return &AbstractChannal{}
}

var memory = String{ len("memory"), "memory" }
var channalMemoryContext = &AbstractContext{
    memory,
    channalContextCreate,
    channalContextInit,
}

func channalContextCreate(cycle *AbstractCycle) unsafe.Pointer {
    return nil
}

func channalContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    return ""
}

var ChannalModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(channalMemoryContext),
    nil,
    CONFIG_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &ChannalModule)
}