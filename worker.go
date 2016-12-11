
/*
 * Copyright (C) 2016 Meng Shi
 */


package main


import (
      "os"
     // "fmt"

    . "worker/types"
    . "worker/define"

    /*
    . "worker/autoconf"
    . "worker/modules"
    */
)


func wkrGetOption(argc int, argv []string) int {
    var i int;

    for i = 1; i < argc; i++ {

	if argv[i][0] != '-' {
	    return WkrError;
	}

        switch argv[i][1] {

        case 'c':
	    if argv[i + 1] == "" {
	        return WkrError;
	    }
/*
	    work.ConfFile.Data = argv[i + 1];
	    workcycle.ConfFile.Len  = len(workcycle.ConfFile.Data.(string)) + 1;
	    */

            i++;

            break;

        case 't':
	    //WkrDvrTestConfig = true;
	    break;

        default:
            break;
        }
    }
/*
    if cycle.ConfFile.Data == nil {
        cycle.ConfFile.Data = DvrConfPath;
        cycle.ConfFile.Len  = len(DvrConfPath) + 1;
    }


    if DvrConfFullName(cycle, &cycle.ConfFile) == DvrError {
	return WkrError;
    }
    */

    return WkrOk;
}


func wkrSetOption(argc int, argv []string) int {
    return WkrOk;
}




func main() {

    argc := len(os.Args)

    if wkrGetOption(argc, os.Args) == WkrError {
        return;
    }

    if wkrSetOption(argc, os.Args) == WkrError {
        return;
    }

}
