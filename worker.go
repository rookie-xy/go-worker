
/*
 * Copyright (C) 2016 Meng Shi
 */


package main


import (
   //   "os"
//      "fmt"

    . "go-worker/types"
    . "go-worker/modules"
    //"fmt"
    "os"
    "fmt"
)


func main() {
    if Modules == nil {
        return
    }

    cycle, err := CoreInit(Modules)
    if err != nil {
        return
    }

    option := cycle.Option

//    option.Create(nil)
    fmt.Println(option.File)

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
