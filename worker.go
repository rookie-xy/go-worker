
/*
 * Copyright (C) 2016 Meng Shi
 */


package main


import (
      "os"
      "fmt"

    . "worker/types"
    . "worker/modules"
)


type worker struct {
    configure  *Configure
    log        *Log
}


func (w *worker) main() {
}


func main() {

    if option := Modules.CoreInit(), ok != nil {
        return
    }

    argc := len(os.Args)

    if option.Get(argc, os.Args) == Error {
        return;
    }

    if option.Set(argc, os.Args) == Error {
        return;
    }

    if cycle := Modules.SystemInit(option), ok != nil {
    }

    Modules.UserInit(cycle)

    for {
        go worker.main()
    }
}
