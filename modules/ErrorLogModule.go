/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "fmt"
    . "go-worker/types"
)

var ErrorLog = String{ len("ErrorLog"), "ErrorLog" }

var ErrorLogCommands = []Command{

    { ErrorLog,
      MainConf|ConfMore1,
      SetErrorLog,
      0,
      0,
      nil },

      NilCommand,
}

var ErrorLogContext = Context{
    ErrorLog,
    nil,
}

var ErrorLogModule = Module{
    0,
    0,
    &ErrorLogContext,
    ErrorLogCommands,
    SYSTEM_MODULE,
    nil,
    nil,
}

func SetErrorLog(cf *AbstractConfigure, cmd *Command, conf interface{}) string {
    fmt.Println("Configure Module Set Error log Finish")
    return ""
}
