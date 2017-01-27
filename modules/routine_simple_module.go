/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
//      "fmt"
//    . "unsafe"

    . "worker/types"
)

type SimpleRoutine struct {
    AbstractRoutine
}

var SimpleRoutineCommands = []Command{
      NilCommand,
}

var RoutineName = String{ len("Routine"), "Routine" }

var SimpleRoutineContext = AbstractContext{
    RoutineName,
    nil,
}

var SimpleRoutineModule = Module{
    0,
    0,
    &SimpleRoutineContext,
    SimpleRoutineCommands,
    CORE_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &SimpleRoutineModule)
}
