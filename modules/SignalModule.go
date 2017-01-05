
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
//      "fmt"
    . "unsafe"

    . "go-worker/types"
)

func init() {
    Modules = append(Modules, &SignalModule)
}

var signal = String{ int(Sizeof("Signal")) - 1, "Signal" }
//var conf = Signal{ "abc" }


var SignalCommands = []Command{
      NilCommand,
}


var SignalContext = Context{
    signal,
    nil,
}


var SignalModule = Module{
    0,
    0,
    &SignalContext,
    SignalCommands,
    CORE_MODULE,
    nil,
    nil,
}

/*
func (s *Signal) Create(cycle *Cycle) {
    fmt.Println("abc")
}


func (s *Signal) Init(cycle *Cycle) {
    fmt.Println("abc")
}
*/


/*
    interface {
        SignalCreate
	SignalInit
    },
    */
