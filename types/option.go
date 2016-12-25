
/*
 * Copyright (C) 2016 Meng Shi
 */


package types

//import "fmt"

type Option struct {
    File  string

    Data  interface {
        Set(argc int, argv []string) int
        Get(argc int, argv []string) int
    }
}

/*
func (o Option) Create(cycle *Cycle) {
    fmt.Println("abc")
}


func (o Option) Init(cycle *Cycle) {
    fmt.Println("abc")
}


func (o Option) Get(argc int, argv []string) int {
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


func (o Option) Set(argc int, argv []string) int {
    return Ok;
}
*/
