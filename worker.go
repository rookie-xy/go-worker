/*
 * Copyright (C) 2016 Meng Shi
 */

package main

import (
      "os"
//      "fmt"
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
        mod := modules[m]

        if mod.Type != CORE_MODULE {
            continue
        }

        if mod.InitModule != nil {
            if mod.InitModule(cycle) == Error {
                os.Exit(2)
            }
        }

        if mod.InitRoutine != nil {
	    if mod.InitRoutine(cycle) == Error {
                os.Exit(2)
            }
        }
    }

    return Ok
}

func (w *worker) SystemInit(configure *AbstractConfigure) int {
    cycle := w.GetCycle()

    if cycle.SetConfigure(configure) == Error {
        return Error
    }

    return Ok
}

func (w *worker) UserInit() int {
    return Ok
}

func main() {
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

    //fmt.Println(worker.cycle.GetOption().GetResult("mengshi"))

    configure := NewConfigure()

    if worker.SystemInit(configure) == Error {
        return
    }

    if worker.UserInit() == Error {
        return
    }

    return
}
