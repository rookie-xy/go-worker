/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
    . "go-worker/types"
)

var SimpleSignalCommands = []Command{
      NilCommand,
}

var SimpleSignalName = String{ len("Signal"), "Signal" }

var SimpleSignalContext = AbstractContext{
    SimpleSignalName,
    nil,
}

var SimpleSignalModule = Module{
    0,
    0,
    &SimpleSignalContext,
    SimpleSignalCommands,
    CORE_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &SimpleSignalModule)
}
