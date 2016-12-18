/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"
    . "unsafe"
    . "worker/types"
)


var Include = String{ DvrSizeof("Include") - 1, "Include" };


var ConfigreCommands = []Command{

    { Include,
      AnyConf|ConfTake1,
      ConfigureInclude,
      0,
      0,
      nil },

      NilCommand,
}


var ConfigureModule = Moudle{
    0,
    0,
    nil,
    ConfigreCommands,
    SYSTEM_MODULE,
    nil,
    nil,
}


func ConfigureInclude(cf *Configure, cmd *Command, conf *Void) string {
    fmt.Println("Configure Module Include Command Finish");
    return ConfigureOk;
}
