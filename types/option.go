/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "fmt"

type Option struct {
    *Log

     argc   int
     argv   []string
     items  map[string]interface{}

     option OptionIf
}

type OptionIf interface {
    Parse() int
}

func NewOption(log *Log) *Option {
    return &Option{
        Log:   log,
        items: make(map[string]interface{}),
    }
}

func (o *Option) GetArgc() int {
    return o.argc
}

func (o *Option) GetArgv() []string {
    return o.argv
}

func (o *Option) SetArgs(argc int, argv []string) int {
    if argc <= 0 || argv == nil {
        return Error
    }

    o.argc = argc
    o.argv = argv

    return Ok
}

func (o *Option) SetItem(k string, v interface{}) {
    o.items[k] = v
}

func (o *Option) GetItem(k string) interface{} {
    return o.items[k]
}

func (o *Option) SetOptionIf(option OptionIf) int {
    if option == nil {
        return Error
    }

    o.option = option

    return Ok
}

func (o *Option) GetOptionIf() OptionIf {
    return o.option
}

func (o *Option) Parse() int {
fmt.Println("option type")
    /*
    log := o.Log

    argv := o.GetArgv()

    for i := 1; i < o.GetArgc(); i++ {

        if argv[i][0] != '-' {
            return Error
        }

        switch argv[i][1] {

        case 'c':
            if argv[i + 1] == "" {
                return Error
            }

            // file://path=/home/
            o.SetItem("configure", "file://resource=" + argv[i + 1])
            i++

            break

        case 'z':
            if argv[i + 1] == "" {
                return Error
            }

            // file://path=/home/
            o.SetItem("configure", "zookeeper://resource=" + argv[i + 1])
            i++

            break

        case 't':
            o.SetItem("test", true)

            break

        default:
            o.SetItem("invaild", "")
            log.Info("not found any option")
            //o.result["error"] = "not found any option"
            break
        }
    }
    */

    return Ok
}
