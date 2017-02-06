/*
 * Copyright (C) 2016 Meng Shi
 */

package types

var (
    Ok    =  0
    Error = -1
    Again = -2
)

const (
    CORE_MODULE = 0x45524F43
    SYSTEM_MODULE = 0x464E4F43
    USER_MODULE = 0x464E4F43
)

type InitFunc func(cycle *AbstractCycle) int
type MainFunc func(cycle *AbstractCycle) int

type Module struct {
    CtxIndex   uint
    Index      uint
    Context   *AbstractContext
    Commands   []Command
    Type       uint
    Init       InitFunc
    Main       MainFunc
}

var Modules = []*Module{}
