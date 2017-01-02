/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "fmt"
    . "go-worker/types"
)

var Include = String{ len("Include"), "Include" }

var ConfigreYamlCommands = []Command{

    { Include,
      AnyConf|ConfTake1,
      ConfigureInclude,
      0,
      0,
      nil },

      NilCommand,
}

var YamlConfigureModule = Module{
    0,
    0,
    nil,
    ConfigreYamlCommands,
    SYSTEM_MODULE,
    nil,
    nil,
}

func ConfigureInclude(cf *Configure, cmd *Command, conf interface{}) string {
    fmt.Println("Configure Module Include Command Finish")
    return ""
}
