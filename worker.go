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
    //_ "github.com/rookie-xy/modules/inputs/httpd/src"

    //_ "github.com/rookie-xy/modules/channels/memory/src"
    _ "github.com/rookie-xy/modules/outputs/stdout/src"
)

type worker struct {
    *Log
    *Cycle
     modules  []*Module
}

func NewWorker(log *Log) *worker {
    return &worker{
        Log:log,
    }
}

func (w *worker) SetModules(m []*Module) int {
    if m == nil {
        return Error
    }

    w.modules = m

    return Ok
}

func (w *worker) GetModules() []*Module {
    return w.modules
}

func (w *worker) SetCycle(cycle *Cycle) int {
    if cycle == nil {
        return Error
    }

    w.Cycle = cycle

    return Ok
}

func (w *worker) GetCycle() *Cycle {
    return w.Cycle
}

func (w *worker) SystemInit(option *Option) int {
    modules, cycle := w.modules, w.Cycle

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

func (w *worker) ConfigureInit(configure *Configure) int {
    cycle := w.GetCycle()

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

    modules := w.modules
    if modules == nil {
        return Error
    }

    configure.Block(CONFIG_MODULE, -1)

    return Ok
}

func (w *worker) Start() int {
    modules, cycle := w.modules, w.Cycle

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

    if cycle := w.Cycle; cycle != nil {
        if cycle.Start() == Error {
            return Error
        }
    }

    return Ok
}

func (w *worker) Stop() int {
    if cycle := w.Cycle; cycle != nil {
        if cycle.Stop() == Error {
            return Error
        }
    }

    return Ok
}

func (w *worker) Monitor() int {
    if cycle := w.Cycle; cycle != nil {
        if routine := cycle.Routine; routine != nil {
            if routine.Monitor() == Error {
                return Error
            }
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

    Modules = append(Modules, nil)
    for /* nil */; Modules[n] != nil; n++ {
        Modules[n].Index = uint(n)
    }

    if n <= 0 {
        log.Info("no module to load")
    }

    worker := NewWorker(log)
    if worker.SetModules(Modules) == Error {
        return
    }

    cycle := NewCycle(log)
    cycle.SetModules(Modules)

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    //cycle := NewCycle(log)
    cycle.Option = option
    worker.Cycle = cycle

    if worker.SystemInit(option) == Error {
        return
    }

    configure := cycle.Configure
    if configure == nil {
        configure = NewConfigure(cycle)
    }

    if worker.ConfigureInit(configure) == Error {
        return
    }

    if worker.Start() == Error {
        return
    }

    select {

    }

    worker.Monitor()

    if worker.Stop() == Error {
        return
    }
    return
}


       /*
        switch module.Type {
        case INPUT_MODULE:
            fmt.Println("input")
        case CHANNEL_MODULE:
            fmt.Println("channel")
        case OUTPUT_MODULE:
            fmt.Println("output")
        case SYSTEM_MODULE:
            fmt.Println("system")
        case CONFIG_MODULE:
            fmt.Println("config")

        default:
            switch module.Type >> 32 {
            case INPUT_MODULE:
                fmt.Println("UUUUUUUUUUUUUUUUUUUUUUU")
            default:
                //fmt.Printf("kkkkkkkkkkkkkkkkkkkkkkkkkkk: %X\n", module.Type>>8)
            }

        }
        */
