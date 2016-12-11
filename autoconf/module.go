
/*
 * Copyright (C) 2016 Meng Shi
 */


package autoconf


import (
    . "devour/types"
    . "devour/modules"
)


var DvrModules = []*DvrMoudleType{
    &DvrCoreModule,
    &DvrConfModule,
    &DvrErrlogModule,
    &DvrNetworkModule,
    &DvrHttpModule,
    //&DvrDefineModule,

    &DvrNetworkCoreModule,
    //&DvrHttpCoreModule,
    //&DvrDefineCoreModule,
    nil,
};
