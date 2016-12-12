
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"
    . "unsafe"

    . "devour/types"
)


var (
    Core            = String{ Sizeof("Core") - 1,            "Core"            };
    User            = String{ Sizeof("User") - 1,            "User"            };
    Daemon          = String{ Sizeof("Daemon") - 1,          "Daemon"          };
    MasterProcess   = String{ Sizeof("Master") - 1,          "Master"          };
    WorkerProcesses = String{ Sizeof("WorkerProcesses") - 1, "WorkerProcesses" };
    PidFile         = String{ Sizeof("Pid") - 1,             "Pid"             };
)


var CoreCommands = []Command{

    { User,
      MainConf|DirectConf|ConfTake12,
      SetUser,
      0,
      0,
      nil },

    { Daemon,
      MainConf|DirectConf|ConfTake1,
      ConfigureSetFlagSlot,
      0,
      Offsetof(Ccf.Daemon),
      nil },

    { MasterProcess,
      MainConf|DirectConf|ConfTake1,
      ConfSetFlagSlot,
      0,
      Offsetof(Ccf.Master),
      nil },

    { WorkerProcesses,
      MainConf|DirectConf|ConfTake1,
      ConfSetNumSlot,
      0,
      Offsetof(Ccf.WorkerProcesses),
      nil },

    { PidFile,
      MainConf|DirectConf|ConfTake1,
      ConfSetStrSlot,
      0,
      Offsetof(Ccf.Pid),
      nil },

      NilCommand,
};


var CoreContext = Context{
    Core,
    CoreModuleCreateConf,
    CoreModuleInitConf,
};


var CoreModule = Moudle{
    0,
    0,
    &CoreContext,
    CoreCommands,
    DVR_CORE_MODULE,
    nil,
    nil,
};


func CoreModuleCreateConf(cycle *CycleType) VoidType {
    ccf := CoreConfType{};
    /*
    if ccf == nil {
        return VoidType(&ConfError)
    }
    */

    ccf.Daemon = ConfUnsetBool;
    ccf.Master = ConfUnsetBool;
    ccf.WorkerProcesses = ConfUnsetInt;

    ccf.User = ConfUnsetInt;
    ccf.Group = ConfUnsetInt;

    fmt.Println("Core Module Create Config Finish");

    return VoidType(&ccf);
}


func CoreModuleInitConf(cycle *CycleType, conf *VoidType) string {
    fmt.Println("Core Module Init Config Finish");
    return ConfOk;
}


func SetUser(cf *Configure, cmd *Command, conf *Void) string {
    value := cf.Args.Elts[1];
    flag := string(value.Data.([]byte));
    fmt.Println("Core Module Set User Finish", flag);
    return ConfOk;
}
