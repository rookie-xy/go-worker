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
    INPUT_CONFIG = 0x00100000

    CHANNAL_MODULE = 0x4749464E4F433
    CHANNAL_CONFIG = 0x00001000

    OUTPUT_MODULE = 0x4749464E4F4345
    OUTPUT_CONFIG = 0x00000010

    CODEC_MODULE = 0x4749464E4F439
    CODEC_CONFIG = 0x01000000

    FILTER_MODULE = 0x4749464E7F439
    FILTER_CONFIG = 0x10000000
)

type InitFunc func(cycle *Cycle) int
type MainFunc func(cycle *Cycle) int

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
