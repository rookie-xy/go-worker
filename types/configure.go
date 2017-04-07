/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

var (
    ConfigOk    =  0
    ConfigError = -1
)

type Configure struct {
    *Cycle
    *Log
    *AbstractFile

     commandType  int
     moduleType   int64
     value        interface{}

     notice       chan int

     Handle
     Parser
}

type Parser interface {
    Marshal(in interface{}) ([]byte, error)
    Unmarshal(in []byte, out interface{}) int
}

func NewConfigure(cycle *Cycle) *Configure {
    return &Configure{
        Cycle        : cycle,
        AbstractFile : NewAbstractFile(cycle.Log),
        notice       : make(chan int),
    }
}

func (c *Configure) SetModuleType(moduleType int64) int {
    if moduleType <= 0 {
        return Error
    }

    c.moduleType = moduleType

    return Ok
}

func (c *Configure) SetCommandType(commandType int) int {
    if commandType <= 0 {
        return Error
    }

    c.commandType = commandType

    return Ok
}

func (c *Configure) SetValue(value interface{}) int {
    if value == nil {
        return Error
    }

    return Ok
}

func (c *Configure) GetValue() interface{} {
    return c.value
}

func (c *Configure) SetNotice(n int) {
    c.notice <- n
}

func (c *Configure) GetNotice() int {
    n := <- c.notice
    return n
}

func (c *Configure) SetHandle(handle Handle) int {
    if handle == nil {
       return Error
    }

    c.Handle = handle

    return Ok
}

func (c *Configure) GetHandle() Handle {
    if c.Handle == nil {
        return nil
    }

    return c.Handle
}

func (c *Configure) SetParser(parser Parser) int {
    if parser == nil {
       return Error
    }

    c.Parser = parser

    return Ok
}

func (c *Configure) GetParser() Parser {
    if c.Parser == nil {
        return nil
    }

    return c.Parser
}

func (c *Configure) Materialized(cycle *Cycle) int {
    if c.value == nil {
        content := c.GetBytes()
        if content == nil {
            /*
            log.Error("configure content: %s, filename: %s, size: %d\n",
                      content, c.GetFileName(), c.GetSize())
                      */
            c.Error("configure content: %s, size: %d\n",
                      content, c.GetSize())

            return Error
        }

        if c.Parser.Unmarshal(content, &c.value) == Error {
            return Error
        }
    }

    switch v := c.value.(type) {

    case []interface{} :
        for _, value := range v {
            c.value = value
            c.Materialized(cycle)
        }

    case map[interface{}]interface{}:
        if c.doParse(v, cycle) == Error {
            return Error
        }

    default:
        c.Warn("unknown")
    }

    return Ok
}

func (c *Configure) doParse(materialized map[interface{}]interface{}, cycle *Cycle) int {
    flag := Ok

    modules := cycle.GetModule(c.moduleType)
    if modules == nil {
        return Error
    }

    for key, value := range materialized {
        if key != nil && value != nil {
            flag = Ok
        }

        name := key.(string)
        found := false

        for m := 0; flag != Error && !found && modules[m] != nil; m++ {
            module := modules[m]

            commands := module.Commands;
            if commands == nil {
                continue;
            }

            for i := 0; commands[i].Name.Len != 0; i++ {

                command := commands[i]

                if len(name) == command.Name.Len &&
                        name == command.Name.Data.(string) {

                				found = true;

                    context := cycle.GetContext(module.Index)

                    c.value = value
																    if cycle.SetConfigure(c) == Error {
                        flag = Error
																				    break
                    }

                    command.Set(cycle, &command, context)
                }
            }
        }
    }

    if flag == Error {
        return ConfigError
    }

    return ConfigOk
}

func (c *Configure)Block(module int64, config int) int {
    var modules []*Module

    if cycle := c.Cycle; cycle == nil {
        return Error
    } else {
        modules = cycle.GetModules()
        if modules == nil || len(modules) <= 0 {
            return Error
        }
    }

    if Block(c.Cycle, modules, module, config) == Error {
        return Error
    }

    return Ok
}

func SetFlag(cycle *Cycle, command *Command, p *unsafe.Pointer) int {
    if cycle == nil || p == nil {
        return Error
    }

    field := (*bool)(unsafe.Pointer(uintptr(*p) + command.Offset))
    if field == nil {
        return Error
    }

    configure := cycle.GetConfigure()
    if configure == nil {
        return Error
    }

    flag := configure.GetValue()
    if flag == true {
        *field = true
    } else if flag == false {
        *field = false
    } else {
        return Error
    }

    /*
    if command.Post != nil {
        post := command.Post.(DvrConfPostType);
        post.Handler(cf, post, *p);
    }
    */

    return Ok
}

func SetString(cycle *Cycle, command *Command, p *unsafe.Pointer) int {
    if cycle == nil || p == nil {
        return Error
    }

    field := (*string)(unsafe.Pointer(uintptr(*p) + command.Offset))
    if field == nil {
        return Error
    }

    configure := cycle.GetConfigure()
    if configure == nil {
        return Error
    }

    strings := configure.GetValue()
    if strings == nil {
        return Error
    }

    *field = strings.(string)

    return Ok
}

func SetNumber(cycle *Cycle, command *Command, p *unsafe.Pointer) int {
    if cycle == nil || p == nil {
        return Error
    }

    field := (*int)(unsafe.Pointer(uintptr(*p) + command.Offset))
    if field == nil {
        return Error
    }

    configure := cycle.GetConfigure()
    if configure == nil {
        return Error
    }

    number := configure.GetValue()
    if number == nil {
        return Error
    }

    *field = number.(int)

    return Error
}

/* default impl */
func (c *Configure) Set() int {
    c.Warn("configure handle set")
    return Ok
}

func (c *Configure) Get() int {
    c.Warn("configure content set")
    return Ok
}

func (c *Configure) Marshal(in interface{}) ([]byte, error) {
    c.Warn("configure Marshal")
    return nil, nil
}

func (c *Configure) Unmarshal(in []byte, out interface{}) int {
    c.Warn("configure Unmarshal")
    return Ok
}
