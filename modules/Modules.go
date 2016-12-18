
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
    . "go-worker/types"
)


var Modules = []*Module{
    &SignalModule,
    &OsModule,
    &RoutineModule,
    &OptionModule,
    &ErrorLogModule,
    &ConfigureYamlModule,
    nil,
}


func CoreInit(modules []*Module) *Option {
    var cycle *Cycle
    var m  int

    for m = 0; modules[m] != nil; m++ {
        mod := modules[m]

        if mod.Type != CORE_MODULE {
            continue
        }

        mod.InitModule(cycle)

        mod.InitRoutine(cycle)
    }

    return nil
}


func SystemInit(option *Option) *Cycle {
    return nil
}


func UserInit(cycle *Cycle) {
}
