/*
 * Copyright (C) 2016 Meng Shi
 */

package main

import (
      "fmt"
      "os"
      "unsafe"
      "strings"

    . "go-worker/types"
    . "go-worker/modules"
)

type worker struct {
    modules  []*Module
    cycle      *AbstractCycle
}

func NewWorker() *worker {
    return &worker{}
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

    w.cycle = cycle

    return Ok
}

func (w *worker) GetCycle() *AbstractCycle {
    return w.cycle
}

func (w *worker) CoreInit(option *AbstractOption) int {
    modules := w.GetModules()
    if modules == nil {
        return Error
    }

    cycle := w.GetCycle()
    if cycle == nil {
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
        } else {
            fmt.Println(this.Name.Data)
        }

        if context := this.Context.Create(cycle); context != nil {
            if *(*string)(unsafe.Pointer(uintptr(context))) == "-1" {

                return Error;
            }

            fmt.Printf("set context: %d\n", module.Index)
            if cycle.SetContext(module.Index, &context) == Error {
                return Error
            }
        }
    }

    config := configure.Get()
    if config == nil {
        return Error
    }

    if config.Parse() == Error {
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
        } else {
            fmt.Println(this.Name.Data)
        }

        fmt.Printf("get context: %d\n", module.Index)
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

func main() {
    Modules = append(Modules, nil)
    for m := 0; Modules[m] != nil; m++ {
        Modules[m].Index++
    }

    fmt.Println(len(Modules))

    worker := NewWorker()
    if worker.SetModules(Modules) == Error {
       return
    }

    option := NewOption()
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    cycle := NewCycle()
    if cycle.SetOption(option) == Error {
        return
    }

    if worker.SetCycle(cycle) == Error {
        return
    }

    if worker.CoreInit(option) == Error {
        return
    }

    configure := NewConfigure()
    if worker.SystemInit(configure) == Error {
        return
    }

    if worker.UserInit() == Error {
        return
    }

    return
}
