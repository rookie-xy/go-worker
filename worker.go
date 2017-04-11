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
    "fmt"
)


func systemInit(cycle *Cycle) int {
    modules:= cycle.GetModules()
    if modules == nil {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]

        if module.Type != SYSTEM_MODULE {
            continue
        }

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

func configureInit(configure *Configure, cycle *Cycle) int {
    if e := <- configure.Event; e != nil {
        opcode := e.GetOpcode()
        if opcode != Ok {
            return Ignore
        }
    }

    if configure.Block(cycle, CONFIG_MODULE, CONFIG_BLOCK) == Error {
        return Error
    }

    return Ok
}

func run(cycle *Cycle) int {
    if cycle.Routine == nil {
        cycle.Routine = NewRoutine()
    }

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

/*
var sigset = [...]int{
    syscall.SIGHUP,
    syscall.SIGTERM,
}
*/

func monitor(cycle *Cycle) int {
    modules := cycle.GetModules()
    if modules == nil {
        return Error
    }

    for {
        select {

        case event := <- cycle.Event:
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
                if Reload(modules) == Error {
                    return Error
                }

            default:
                fmt.Println("ddddddddddddddddddd")
            }

        default:
            //fmt.Println("aaaaaaaaaaaaaaaaa")
        }
    }

    if routine := cycle.Routine; routine != nil {
        if routine.Monitor() == Error {
            return Error
        }
    }

//    signalChan := make(chan os.Signal, 1)
//    signal.Notify(signalChan, sigset)

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

    if configureInit(configure, cycle) == Error {
        return
    }

    if run(cycle) == Error {
        return
    }

    monitor(cycle)

    if stop(cycle) == Error {
        return
    }

    return
}
