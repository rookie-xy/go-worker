
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


func (m []*Module) CoreInit() *Option {
    for i := 0; i < len; m[i] {

    }
}


func (m []*Module) SystemInit(option *Option) *Cycle {
}


func (m []*Module) UserInit(cycle *Cycle) {
}
