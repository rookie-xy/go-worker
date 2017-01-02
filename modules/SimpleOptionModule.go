/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
    . "go-worker/types"
)

type SimpleOption struct {
    *AbstractOption
}

var SimpleOptionModule = Module{
    0,
    0,
    nil,
    nil,
    CORE_MODULE,
    InitModule,
    nil,
}

func InitModule(cycle *AbstractCycle) int {
    simpleOption := &SimpleOption{ cycle.GetOption() }

    if simpleOption.Parse() == Error {
        return Error
    }

    return Ok
}

func (o *SimpleOption) Parse() int {
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

            o.SetResult("config", argv[i + 1])
            i++
            break

        case 't':
            o.SetResult("test", true)
	    break

        default:
            o.SetResult("invaild", "")
            //o.result["error"] = "not found any option"
            break
        }
    }

    o.SetResult("mengshi", "zhangyue")

    return Ok
}
