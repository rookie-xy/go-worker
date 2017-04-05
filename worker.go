/*
 * Copyright (C) 2016 Meng Shi
 */

package main

import (
    "os"
    "strings"

    . "github.com/rookie-xy/worker/types"
    _ "github.com/rookie-xy/worker/modules"

    _ "github.com/rookie-xy/modules/option/simple/src"
    _ "github.com/rookie-xy/modules/configure/yaml/src"
    //_ "github.com/rookie-xy/modules/logs/mlog/src"

    _ "github.com/rookie-xy/modules/inputs/stdin/src"
    _ "github.com/rookie-xy/modules/inputs/httpd/src"

    _ "github.com/rookie-xy/modules/channels/memory/src"
    _ "github.com/rookie-xy/modules/outputs/stdout/src"
)


func systemInit(cycle *Cycle) int {
    modules, option := cycle.GetModules(), cycle.GetOption()

    if modules == nil || cycle == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]



        if module.Type != SYSTEM_MODULE {
            continue
        }

        if module.Init != nil {
            if module.Init(cycle) == Error {
                os.Exit(2)
            }
        }

        if module.Main != nil {
	           if module.Main(cycle) == Error {
                os.Exit(2)
            }
        }
    }

    if option.Materialized() == Error {
        return Error
    }

    return Ok
}

func configureInit(configure *Configure) int {
    cycle := configure.Cycle

    option := cycle.GetOption()
    if option == nil {
        return Error
    }

    item := option.GetItem("configure")
    if item == nil {
        return Error
    }

    file := item.(string)

    fileType := file[0 : strings.Index(file, ":")]
    if fileType == "" {
        return Error
    }

    if configure.SetFileType(fileType) == Error {
        return Error
    }

    fileName := file[strings.LastIndex(file, "/") + 1 : ]
    if fileName == "" {
        return Error
    }

    if configure.SetFileName(fileName) == Error {
        return Error
    }

    // TODO
    resource := file[strings.Index(file, "=") + 1 : ]
    if resource == "" {
        return Error
    }

    if configure.SetResource(resource) == Error {
        return Error
    }

    if cycle.SetConfigure(configure) == Error {
        return Error
    }

    if configure.Block(CONFIG_MODULE, CONFIG_BLOCK) == Error {
        return Error
    }

    return Ok
}

func start(cycle *Cycle) int {
    modules := cycle.GetModules()

    if modules == nil && cycle == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]

        if init := module.Init; init != nil {
            if init(cycle) == Error {
                return Error
            }
        }
    }

    if cycle.Start() == Error {
        return Error
    }

    return Ok
}

func stop(cycle *Cycle) int {
    if cycle.Stop() == Error {
        return Error
    }

    return Ok
}

func monitor(cycle *Cycle) int {
    if routine := cycle.Routine; routine != nil {
        if routine.Monitor() == Error {
            return Error
        }
    }

    return Ok
}

func main() {
    n   := 0
    log := NewLog()

    if log.Set(log) == Error {
        return
    }

    Modules = Load(Modules, nil)
    for /* nil */; Modules[n] != nil; n++ {
        Modules[n].Index = uint(n)
    }

    if n <= 0 {
        log.Info("no module to load")
    }

    cycle := NewCycle(log)
    cycle.SetModules(Modules)

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    cycle.Option = option

    if systemInit(cycle) == Error {
        return
    }

    configure := cycle.Configure
    if configure == nil {
        configure = NewConfigure(cycle)
    }

    if configureInit(configure) == Error {
        return
    }

    if start(cycle) == Error {
        return
    }

    select {

    }

    monitor(cycle)

    if stop(cycle) == Error {
        return
    }

    return
}
