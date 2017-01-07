/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
    //  "fmt"
    //. "unsafe"

    . "go-worker/types"
)

var SimpleOsCommands = []Command{
      NilCommand,
}

var SimpleOsName = String{ len("os"), "os" }

var SimpleOsContext = AbstractContext{
    SimpleOsName,
    nil,
}

var SimpleOsModule = Module{
    0,
    0,
    &SimpleOsContext,
    SimpleOsCommands,
    CORE_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &SimpleOsModule)
}
