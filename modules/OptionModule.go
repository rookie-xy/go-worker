
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"
    . "unsafe"

    . "go-worker/types"
)


var OptionCommands = []Command{
      NilCommand
};


var OptionContext = Context{
    Core,
    Conf : interface {
        Create,
	Init
    }
};


var OptionModule = Moudle{
    0,
    0,
    &OptionContext,
    OptionCommands,
    CORE_MODULE,
    nil,
    nil,
};


func Get(argc int, argv []string) int {
    var i int;

    for i = 1; i < argc; i++ {

	if argv[i][0] != '-' {
	    return Error;
	}

        switch argv[i][1] {

        case 'c':
	    if argv[i + 1] == "" {
	        return Error;
	    }

            i++;

            break;

        case 't':
	    break;

        default:
            break;
        }
    }

    return Ok;
}


func Set(argc int, argv []string) int {
    return Ok;
}
