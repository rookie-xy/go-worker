
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"

    . "go-worker/types"
)


type OptionDefault Option


var option = String{ len("Option"), "Option" }
var optionConf = OptionDefault{ "Option" }


var OptionCommands = []Command{
      NilCommand,
}


var OptionContext = Context{
    option,
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


func (od OptionDefault) Create(cycle *Cycle) {
    fmt.Println("abc")
}


func (od OptionDefault) Init(cycle *Cycle) {
    fmt.Println("abc")
}



func Get(argc int, argv []string) int {
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


func Set(argc int, argv []string) int {
    return Ok;
}
