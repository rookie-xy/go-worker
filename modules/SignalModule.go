
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"
    . "unsafe"

    . "devour/types"
)


var SingalCommands = []Command{
      NilCommand
};


var SingalContext = Context{
    Core,
    Conf : interface {
        Create,
	Init
    }
};


var SingalModule = Moudle{
    0,
    0,
    &SingalContext,
    SingalCommands,
    CORE_MODULE,
    nil,
    nil,
};


func Create(cycle *Cycle) {
}


func Init(cycle *Cycle) {
}
