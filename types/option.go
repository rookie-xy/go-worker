/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractOption struct {
    *AbstractLog
     argc   int
     argv   []string
     items  map[string]interface{}
}

type Option interface {
    Parse() int
}

func NewOption(log *AbstractLog) *AbstractOption {
    return &AbstractOption{
        AbstractLog:log,
        items: make(map[string]interface{}),
    }
}

func (o *AbstractOption) GetArgc() int {
    return o.argc
}

func (o *AbstractOption) GetArgv() []string {
    return o.argv
}

func (o *AbstractOption) SetArgs(argc int, argv []string) int {
    if argc <= 0 || argv == nil {
        return Error
    }

    o.argc = argc
    o.argv = argv

    return Ok
}

func (o *AbstractOption) SetItem(k string, v interface{}) {
    o.items[k] = v
}

func (o *AbstractOption) GetItem(k string) interface{} {
    return o.items[k]
}

func (o *AbstractOption) Parse() int {
	log := o.AbstractLog

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

func initOptionModule(cycle *AbstractCycle) int {
	/*
	simpleOption := NewOption(cycle.log)

	option := cycle.GetOption()
	if option == nil {
		return Error
	}

	log := option.AbstractLog

	if simpleOption.SetOption(option) == Error {
		log.Error("set option error")
		return Error
	}

	if simpleOption.Parse() == Error {
		return Error
	}
	*/

	return Ok
}

var optionModule = Module{
	MODULE_V1,
	CONTEXT_V1,
	nil,
	nil,
	SYSTEM_MODULE,
	initOptionModule,
	nil,
}

func init() {
	Modules = append(Modules, &optionModule)
}
