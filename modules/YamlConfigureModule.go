/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "fmt"
    . "go-worker/types"
)

func init() {
    Modules = append(Modules, &YamlConfigureModule)
}

type YamlConfigure struct {
    *AbstractConfigure
}

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

func ConfigureInclude(cf *AbstractConfigure, cmd *Command, conf interface{}) string {
    fmt.Println("Configure Module Include Command Finish")
    return ""
}
