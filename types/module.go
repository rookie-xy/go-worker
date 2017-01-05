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

type InitModuleFunc func(cycle *AbstractCycle) int
type InitRoutineFunc func(cycle *AbstractCycle) int

type Module struct {
    CtxIndex      uint
    Index         uint
    Context      *Context
    Commands      []Command
    Type          uint
    InitModule    InitModuleFunc
    InitRoutine   InitRoutineFunc
}

//var Modules = []*Module{}
