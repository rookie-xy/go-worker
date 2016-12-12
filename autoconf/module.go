
/*
 * Copyright (C) 2016 Meng Shi
 */


package autoconf


import (
    . "devour/types"
    . "devour/modules"
)


var Modules = []*Moudle{
    &CoreModule,
    &LogModule,
    &ConfigureYamlModule,
    nil,
}
