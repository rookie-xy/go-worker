/*
 * Copyright (C) 2016 Meng Shi
 */

package main

import (
      "fmt"
      "os"
      "unsafe"
      "strings"

    . "worker/types"
    . "worker/modules"
)

type worker struct {
    *AbstractLog
    *AbstractCycle
     modules  []*Module
}

func NewWorker(log *AbstractLog) *worker {
    return &worker{
        AbstractLog:log,
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

func (w *worker) SetCycle(cycle *AbstractCycle) int {
    if cycle == nil {
        return Error
    }

    w.AbstractCycle = cycle

    return Ok
}

func (w *worker) GetCycle() *AbstractCycle {
    return w.AbstractCycle
}

func (w *worker) CoreInit(option *AbstractOption) int {
    modules, cycle := w.modules, w.AbstractCycle

    if modules == nil || cycle == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]

        if module.Type != CORE_MODULE {
            continue
        }

        if module.InitModule != nil {
            if module.InitModule(cycle) == Error {
                os.Exit(2)
            }
        }

        if module.InitRoutine != nil {
	    if module.InitRoutine(cycle) == Error {
                os.Exit(2)
            }
        }
    }

    return Ok
}

func (w *worker) SystemInit(configure *AbstractConfigure) int {
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

    modules := w.GetModules()
    if modules == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]
        if module.Type != SYSTEM_MODULE {
            continue
        }

	this := module.Context
        if this == nil || this.Context == nil {
            continue
        }/* else {
            fmt.Println(this.Name.Data)
        }*/

        if context := this.Context.Create(cycle); context != nil {
            if *(*string)(unsafe.Pointer(uintptr(context))) == "-1" {
                return Error;
            }

            if cycle.SetContext(module.Index, &context) == Error {
                return Error
            }
        }
    }

    if configure.Parse() == Error {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]
        if module.Type != SYSTEM_MODULE {
            continue
        }

	this := module.Context
        if this == nil || this.Context == nil {
            continue
        }/* else {
            fmt.Println(this.Name.Data)
        }*/

        context := cycle.GetContext(module.Index)
        if context == nil {
            continue
        }

        if this.Context.Init(cycle, context) == "-1" {
            return Error
        }
    }

    return Ok
}

func (w *worker) UserInit() int {
    return Ok
}

func (w *worker) Start() int {
    // TODO other need init

    if cycle := w.AbstractCycle; cycle != nil {
        if cycle.Start() == Error {
            return Error
        }
    }

    return Ok
}

func (w *worker) Stop() int {
    if cycle := w.AbstractCycle; cycle != nil {
        if cycle.Stop() == Error {
            return Error
        }
    }

    return Ok
}

func (w *worker) Monitor() int {
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
        Modules[n].Index++
    }

    fmt.Println(len(Modules), MODULE_VER)

    if n <= 0 {
        log.Info("no module to load")
    }

    worker := NewWorker(log)
    if worker.SetModules(Modules) == Error {
        return
    }

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    cycle := NewCycle(log)
    cycle.AbstractOption = option
    worker.AbstractCycle = cycle

    if worker.CoreInit(option) == Error {
        return
    }

    configure := NewConfigure(log)
    if worker.SystemInit(configure) == Error {
        return
    }

    if worker.UserInit() == Error {
        return
    }

    if worker.Start() == Error {
        return
    }

    worker.Monitor()

    if worker.Stop() == Error {
        return
    }

    return
}
