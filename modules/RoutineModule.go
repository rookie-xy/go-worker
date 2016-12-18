
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"
    . "unsafe"

    . "go-worker/types"
)


var RoutineCommands = []Command{
      NilCommand
};


var RoutineContext = Context{
    Core,
    Conf : interface {
        Create,
	Init
    }
};


var RoutineModule = Moudle{
    0,
    0,
    &RoutineContext,
    RoutineCommands,
    CORE_MODULE,
    nil,
    nil,
};


func Create(cycle *Cycle) {
}


func Init(cycle *Cycle) {
}
