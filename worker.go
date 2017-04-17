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


func systemInit(cycle *Cycle) int {
    modules:= cycle.GetSomeModules(SYSTEM_MODULE)
    if modules == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]

        if module.Init != nil {
            if module.Init(cycle) == Error {
                os.Exit(1)
            }
        }

        if main := module.Main; main != nil {
	           if main.Start(cycle) == Error {
                os.Exit(2)
            }
        }
    }

    return Ok
}

func configureInit(c *Cycle) int {
    /*
    if e := <- configure.Event; e != nil {
        opcode := e.GetOpcode()
        if opcode != Ok {
            return Ignore
        }
    }
    */

    select {

    case e := <- c.Event:
        if op := e.GetOpcode(); op != LOAD {
            return Ignore
        }
    }

    if c.Block(c, CONFIG_MODULE, CONFIG_BLOCK) == Error {
        return Error
    }

    return Ok
}

func run(cycle *Cycle) int {
    if cycle.Routine == nil {
        cycle.Routine = NewRoutine()
    }

    modules := cycle.GetSpacModules()
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

    if cycle.Start(modules) == Error {
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

/*
var sigset = [...]int{
    syscall.SIGHUP,
    syscall.SIGTERM,
}
*/

func monitor(c *Cycle) int {
    modules := c.GetModules()
    if modules == nil {
        return Error
    }

    for {
        select {

        case event := <- c.GetNotice():
            opcode := event.GetOpcode()

            switch opcode {

            case START:
                if Start(modules) == Error {
                    return Error
                }

            case STOP:
                if Stop(modules) == Error {
                    return Error
                }

            case RELOAD:
                if c.Reload() == Error {
                    return Error
                }
            }
        }
    }

    if routine := c.Routine; routine != nil {
        if routine.Monitor() == Error {
            return Error
        }
    }

//    signalChan := make(chan os.Signal, 1)
//    signal.Notify(signalChan, sigset)

    return Ok
}

func exit(cycle *Cycle) {
    return
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

    if configureInit(cycle) == Error {
        return
    }

    if run(cycle) == Error {
        return
    }

    monitor(cycle)

    exit(cycle)

    return
}
