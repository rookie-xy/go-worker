/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"
    . "devour/types"
)


var ErrorLog = String{ Sizeof("ErrorLog") - 1, "ErrorLog" };


var ErrorLogCommands = []Command{

    { ErrorLog,
      MainConf|ConfMore1,
      SetErrorLog,
      0,
      0,
      nil },

      NilCommand,
};


var ErrorLogContext = Context{
    ErrorLog,
    nil,
    nil,
};


var ErrorLogModule = Moudle{
    0,
    0,
    &ErrorLogContext,
    ErrorLogCommands,
    DVR_CORE_MODULE,
    nil,
    nil,
};


func SetErrorLog(cf *Configure, cmd *Command, conf *Void) string {
    fmt.Println("Configure Module Set Error log Finish");
    return ConfigureOk;
}
