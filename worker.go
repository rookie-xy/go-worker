/*
 * Copyright (C) 2016 Meng Shi
 */

package main

import (
      "os"

    . "github.com/rookie-xy/worker/types"

    _ "github.com/rookie-xy/worker/modules"
    _ "github.com/rookie-xy/modules/option/simple/src"
    _ "github.com/rookie-xy/modules/configure/file/src"
    _ "github.com/rookie-xy/modules/configure/yaml/src"
    //_ "github.com/rookie-xy/modules/logs/mlog/src"
    _ "github.com/rookie-xy/modules/inputs/stdin/src"
    _ "github.com/rookie-xy/modules/inputs/httpd/src"
    _ "github.com/rookie-xy/modules/channels/memory/src"
    _ "github.com/rookie-xy/modules/outputs/stdout/src"
)

/*
var sigset = [...]int{
    syscall.SIGHUP,
    syscall.SIGTERM,
}

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, sigset)
*/

func main() {
    count := 0
    log   := NewLog()

    Modules = Load(Modules, nil)
    for /* nil */; Modules[count] != nil; count++ {
        Modules[count].Index = uint(count)
    }

    if count <= 0 {
        log.Info("no module to load")
    }

    cycle := NewCycle(log)

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    cycle.Option = option

    cycle.Init(Modules)

    cycle.Main(Modules)

    cycle.Monitor(Modules)

    cycle.Exit(Modules)

    return
}
