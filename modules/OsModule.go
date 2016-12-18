
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"
    . "unsafe"

    . "go-worker/types"
)


var OsCommands = []Command{
      NilCommand
};


var OsContext = Context{
    Core,
    Conf : interface {
        Create,
	Init
    }
};


var OsModule = Moudle{
    0,
    0,
    &OsContext,
    OsCommands,
    CORE_MODULE,
    nil,
    nil,
};


func Create(cycle *Cycle) {
}


func Init(cycle *Cycle) {
}
