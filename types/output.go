/*
 * Copyright (C) 2017 Meng Shi
 */

package types

import (
    "unsafe"
)

type AbstractOutput struct {
    *AbstractCycle

     name   string
     output  Output
}

type Output interface {}

func NewOutput() *AbstractOutput {
    return &AbstractOutput{}
}

var stdout = String{ len("stdout"), "stdout" }
var outputStdoutContext = &AbstractContext{
    stdout,
    outputContextCreate,
    outputContextInit,
}

func outputContextCreate(cycle *AbstractCycle) unsafe.Pointer {
    return nil
}

func outputContextInit(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    return ""
}

var OutputModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(outputStdoutContext),
    nil,
    CONFIG_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &OutputModule)
}
