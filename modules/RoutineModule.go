
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
//      "fmt"
//    . "unsafe"

    . "go-worker/types"
)


type RoutineImpl Routine


var routine = String{ len("Routine"), "Routine" }
var routineConf = RoutineImpl{ "Option" }


var RoutineCommands = []Command{
      NilCommand,
}


var RoutineContext = Context{
    routine,
    routineConf,
}


var RoutineModule = Module{
    0,
    0,
    &RoutineContext,
    RoutineCommands,
    CORE_MODULE,
    nil,
    nil,
}


func (ri RoutineImpl) Create(cycle *Cycle) {
}


func (ri RoutineImpl) Init(cycle *Cycle) {
}
