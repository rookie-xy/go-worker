/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
    . "go-worker/types"
)

var ErrorLogName = String{ len("ErrorLog"), "ErrorLog" }

var ErrorLogContext = AbstractContext{
    ErrorLogName,
    nil,
}

var ErrorLogModule = Module{
    0,
    0,
    &ErrorLogContext,
    nil,
    SYSTEM_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &ErrorLogModule)
}
