
/*
 * Copyright (C) 2016 Meng Shi
 */


package autoconf


import (
    . "devour/types"
)


var Modules = []*Moudle{
    &SignalModule,
    &OsModule,
    &RoutineModule,
    &OptionModule,
    &LogModule,
    &ConfigureYamlModule,
    nil,
}


func (modules []*Module) CoreInit() *Option {
    var m  int

    for m = 0; modules[m] != nil; m++ {

        if modules[m].Type != CORE {
            continue
        }

        modules[m].InitModule(cycle)
        modules[m].InitRoutine(cycle)
    }
}


func (modules []*Module) SystemInit(option *Option) *Cycle {
}


func (modules []*Module) UserInit(cycle *Cycle) {
}
