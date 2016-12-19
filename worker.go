
/*
 * Copyright (C) 2016 Meng Shi
 */


package main


import (
      "os"
//      "fmt"

    . "go-worker/types"
    . "go-worker/modules"
)


func main() {
    //m := len(Modules)

    option, err := Modules[0].CoreInit(Modules)
    if err != nil {
        return
    }

    argc := len(os.Args)

    if option.Data.Get(argc, os.Args) == Error {
        return
    }

    if option.Data.Set(argc, os.Args) == Error {
        return
    }

/*
    if cycle := Modules[m].SystemInit(option), ok != nil {
    }

    Modules[m].UserInit(cycle)

    for {
        go worker.main()
    }
*/

    return
}
