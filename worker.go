/*
 * Copyright (C) 2016 Meng Shi
 */

package main

import (
      "os"
      "unsafe"
      "strings"

    . "github.com/rookie-xy/worker/types"

    _ "github.com/rookie-xy/worker-input-modules/stdin_modules/src"
    _ "github.com/rookie-xy/worker-input-modules/http_modules/src"
    _ "github.com/rookie-xy/worker-channel-modules/memory_modules/src"
    _ "github.com/rookie-xy/worker-output-modules/stdout_modules/src"
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

func (w *worker) SystemInit(option *AbstractOption) int {
    modules, cycle := w.modules, w.AbstractCycle

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

    return Ok
}

func (w *worker) ConfigureInit(configure *AbstractConfigure) int {
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

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]
        if module.Type != CONFIG_MODULE {
            continue
        }

        context := (*AbstractContext)(unsafe.Pointer(module.Context))
        if context == nil {

            continue
        }

        if handle := context.Create; handle != nil {
            this := handle(cycle)

            //fmt.Println(module.Index)
            if cycle.SetContext(module.Index, &this) == Error {
                return Error
            }
        }
    }

    if configure.Parse(cycle) == Error {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]
        if module.Type != CONFIG_MODULE {
            continue
        }

        this := (*AbstractContext)(unsafe.Pointer(module.Context))
        if this == nil {
            continue
        }

        context := cycle.GetContext(module.Index)
        if context == nil {
            continue
        }

        if init := this.Init; init != nil {
            if init(cycle, context) == "-1" {
                return Error
            }
        }
    }

    return Ok
}

func (w *worker) Start() int {
    modules, cycle := w.modules, w.AbstractCycle

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
    if cycle := w.AbstractCycle; cycle != nil {
        if routine := cycle.AbstractRoutine; routine != nil {
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

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    cycle := NewCycle(log)
    cycle.AbstractOption = option
    worker.AbstractCycle = cycle

    if worker.SystemInit(option) == Error {
        return
    }


    configure := NewConfigure(log)
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
