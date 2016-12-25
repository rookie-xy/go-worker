
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


import (
    "os"
)


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


type InitModuleFunc func(cycle *Cycle) int
type InitRoutineFunc func(cycle *Cycle) int


type Module struct {
    CtxIndex      uint
    Index         uint
    Context      *Context
    Commands      []Command
    Type          uint
    InitModule    InitModuleFunc
    InitRoutine   InitRoutineFunc
}


func CoreInit(modules []*Module) (*Cycle, error) {
    var cycle *Cycle
    var m  int

    cycle = &Cycle{}

    for m = 0; modules[m] != nil; m++ {
        mod := modules[m]

        if mod.Type != CORE_MODULE {
            continue
        }

        if mod.InitModule != nil {
            if mod.InitModule(cycle) == Error {
                os.Exit(2)
            }
        }

        if mod.InitRoutine != nil {
	    if mod.InitRoutine(cycle) == Error {
                os.Exit(2)
            }
        }
    }

    return cycle, nil
}


func SystemInit(option *Option) *Cycle {
    return nil
}


func UserInit(cycle *Cycle) {
}
