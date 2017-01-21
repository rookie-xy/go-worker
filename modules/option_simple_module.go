/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
    . "go-worker/types"
)

type simpleOption struct {
    *AbstractOption
}

func NewSimpleOption() *simpleOption {
    return &simpleOption{ NewOption(nil) }
}

func (so *simpleOption) SetOption(option *AbstractOption) int {
    if option == nil {
        return Error
    }

    so.AbstractOption = option

    return Ok
}

func (so *simpleOption) GetOption() *AbstractOption {
    return so.AbstractOption
}

func (o *simpleOption) Parse() int {
    log := o.AbstractLog

    argv := o.GetArgv()

    for i := 1; i < o.GetArgc(); i++ {

	if argv[i][0] != '-' {
	    return Error
	}

        switch argv[i][1] {

        case 'c':
	    if argv[i + 1] == "" {
	        return Error
	    }

            log.Error("aaaaaaaaaaaaa")

            // file://path=/home/
            o.SetItem("configure", "file://path=" + argv[i + 1])
            i++
            break

        case 'z':
	    if argv[i + 1] == "" {
	        return Error
	    }

            // file://path=/home/
            o.SetItem("configure", "zookeeper://path=" + argv[i + 1])
            i++
            break

        case 't':
            o.SetItem("test", true)
	    break

        default:
            o.SetItem("invaild", "")
            //o.result["error"] = "not found any option"
            break
        }
    }

    return Ok
}

func initSimpleOptionModule(cycle *AbstractCycle) int {
    simpleOption := NewSimpleOption()

    option := cycle.GetOption()
    if option == nil {
        return Error
    }

    log := option.AbstractLog

    if simpleOption.SetOption(option) == Error {
        log.Error("set option error")
        return Error
    }

    if simpleOption.Parse() == Error {
        return Error
    }

    return Ok
}

var SimpleOptionModule = Module{
    0,
    0,
    nil,
    nil,
    CORE_MODULE,
    initSimpleOptionModule,
    nil,
}

func init() {
    Modules = append(Modules, &SimpleOptionModule)
}
