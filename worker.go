
/*
 * Copyright (C) 2016 Meng Shi
 */


package main


import (
      "os"
//      "fmt"

    . "go-worker/types"
    . "go-worker/modules"
    "fmt"
)


func main() {
    //m := len(Modules)
fmt.Println("init finish")
    option := CoreInit(Modules)
    if option != nil {
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
