/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

var (
    Ok    =  0
    Error = -1
    Again = -2
)

const (
    MODULE_V1 = 0
    CONTEXT_V1 = 0

    SYSTEM_MODULE = 0x4D4554535953
    CONFIG_MODULE = 0x4749464E4F43

    INPUT_MODULE = 0x4749464E4F43234
    CHANNAL_MODULE = 0x4749464E4F433
    OUTPUT_MODULE = 0x4749464E4F4345
)

type InitFunc func(cycle *AbstractCycle) int
type MainFunc func(cycle *AbstractCycle) int

type Module struct {
    CtxIndex   uint
    Index      uint
    Context    unsafe.Pointer
    Commands   []Command
    Type       int64
    Init       InitFunc
    Main       MainFunc
}

var Modules = []*Module{}
