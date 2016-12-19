
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"

    . "go-worker/types"
)


type option Option


var optionName = String{ len("Option"), "Option" }
var optionConf = option{ "Option", option{} }


var OptionCommands = []Command{
      NilCommand,
}


var OptionContext = Context{
    optionName,
    optionConf,
}


var OptionModule = Module{
    0,
    0,
    &OptionContext,
    OptionCommands,
    CORE_MODULE,
    nil,
    nil,
}


func (o option) Create(cycle *Cycle) {
    fmt.Println("abc")
}


func (o option) Init(cycle *Cycle) {
    fmt.Println("abc")
}


func (o option) Get(argc int, argv []string) int {
    var i int

    for i = 1; i < argc; i++ {

	if argv[i][0] != '-' {
	    return Error
	}

        switch argv[i][1] {

        case 'c':
	    if argv[i + 1] == "" {
	        return Error
	    }

            i++

            break

        case 't':
	    break

        default:
            break
        }
    }

    return Ok
}


func (o option) Set(argc int, argv []string) int {
    return Ok;
}
