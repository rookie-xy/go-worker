/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type Option struct {
    *Log

     argc   int
     argv   []string
     items  map[string]interface{}
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

func (o *Option) Parse() int {
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

    return Ok
}
