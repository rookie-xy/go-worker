/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
    . "github.com/rookie-xy/worker/types"
)

func initOptionModule(cycle *Cycle) int {
    option := cycle.GetOption()
    if option == nil {
        return Error
    }

    if option.Parse() == Error {
        return Error
    }

    return Ok
}

var optionModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    nil,
    SYSTEM_MODULE,
    initOptionModule,
    nil,
}

func init() {
    Modules = append(Modules, &optionModule)
}