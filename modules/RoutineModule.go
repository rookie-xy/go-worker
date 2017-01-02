/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
//      "fmt"
//    . "unsafe"

    . "go-worker/types"
)

type SimpleRoutine struct {
    AbstractRoutine
}
/*
var routine = String{ len("Routine"), "Routine" }
var routineConf = SimpleRoutine{ AbstractRoutine: AbstractRoutine{} }

var RoutineCommands = []Command{
      NilCommand,
}

var RoutineContext = Context{
    routine,
    routineConf,
}
*/

var RoutineModule = Module{
    0,
    0,
    /*
    &RoutineContext,
    RoutineCommands,
    */
    nil,
    nil,
    CORE_MODULE,
    nil,
    nil,
}

func (sr SimpleRoutine) Create(cycle *Cycle) {
}

func (sr SimpleRoutine) Init(cycle *Cycle) {
}
