/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
	"fmt"
//    . "unsafe"

    . "devour/types"
)


var DvrInclude = DvrStrType{ DvrSizeof("Include") - 1, "Include" };


var DvrConfCommands = []DvrCommandType{

    { DvrInclude,
      DvrAnyConf|DvrConfTake1,
      DvrConfInclude,
      0,
      0,
      nil },

      DvrNilCommand,
}


var DvrConfModule = DvrMoudleType{
    0,
    0,
    nil,
    DvrConfCommands,
    DVR_CONF_MODULE,
    nil,
    nil,
}


func DvrConfInclude(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
    fmt.Println("Conf Module Conf Include Finish");
    return DvrConfOk;
}
