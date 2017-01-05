/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
    //  "fmt"
    //. "unsafe"

    . "go-worker/types"
)

func init() {
    Modules = append(Modules, &OsModule)
}

var os = String{ len("Os"), "Os" }

var OsCommands = []Command{
      NilCommand,
};

var OsContext = Context{
    os,
    nil,
};

var OsModule = Module{
    0,
    0,
    &OsContext,
    OsCommands,
    CORE_MODULE,
    nil,
    nil,
};

/*
func Create(cycle *Cycle) {
}


func Init(cycle *Cycle) {
}
*/
