
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
    &SimpleOptionModule,
    &ErrorLogModule,
    &YamlConfigureModule,
    nil,
}
