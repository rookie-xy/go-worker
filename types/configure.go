/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "fmt"
    "unsafe"
)

var (
    ConfigOk    =  0
    ConfigError = -1
)

type Configure struct {
    *Cycle
    *Log
    *File

     resource     string
     fileName     string
     commandType  int
     moduleType   int64
     value        interface{}
     Content
     //Parser
     configure    ConfigureIf
}

type Content interface {
    Set() int
    Get() int
}
/*
type Parser interface {
    Marshal(in interface{}) ([]byte, int)
    Unmarshal(in []byte, out interface{}) int
}
*/

type ConfigureIf interface {
    Parser(in []byte, out interface{}) int
}

func NewConfigure(cycle *Cycle) *Configure {
    return &Configure{
        Cycle: cycle,
        File : NewFile(cycle.Log),
    }
}

func (c *Configure) SetName(file string) int {
    if file == "" {
        return Error
    }

    if c.File.SetName(file) == Error {
        return Error
    }

    return Ok
}

func (c *Configure) GetName() string {
    return c.File.GetName()
}

func (c *Configure) SetFileName(fileName string) int {
    if fileName == "" {
        return Error
    }

    c.fileName = fileName

    return Ok
}

func (c *Configure) GetFileName() string {
    return c.fileName
}

func (c *Configure) SetFileType(fileType string) int {
    if fileType == "" {
        return Error
    }

    if c.SetName(fileType) == Error {
        return Error
    }

    return Ok
}

func (c *Configure) GetFileType() string {
    if fileType := c.GetName(); fileType != "" {
        return fileType
    }

    return ""
}

func (c *Configure) SetFile(action IO) int {
    if action == nil {
        return Error
    }

    if c.File.Set(action) == Error {
        return Error
    }

    return Ok
}

func (c *Configure) GetFile() IO {
    if file := c.File.Get(); file != nil {
        return file
    }

    return nil
}

func (c *Configure) SetResource(resource string) int {
    if resource == "" {
        return Error
    }

    c.resource = resource

    return Ok
}

func (c *Configure) GetResource(resource string) string {
    return c.resource
}

func (c *Configure) Get() ConfigureIf {
    log := c.Log

    file := c.File.Get()
    if file == nil {
        file = NewFile(c.Log)
    }

    if file.Open(c.resource) == Error {
        log.Error("configure open file error")
        return nil
    }

    if file.Read() == Error {
        log.Error("configure read file error")
        goto JMP_CLOSE
        return nil
    }

    if content := file.Type().GetContent(); content != nil {
        c.content = content
    } else {
        log.Warn("not found content: %d\n", 10)
    }

JMP_CLOSE:
    if file.Close() == Error {
        log.Warn("file close error: %d\n", 10)
        return nil
    }

    return c.configure
}

func (c *Configure) Set(configre ConfigureIf) int {
    if configre == nil {
        return Error
    }

    c.configure = configre

    return Ok
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

func (c *Configure) GetValue() interface{} {
    return c.value
}
/*
func (c *Configure) SetContent(content Content) int {
    if content == nil {
       return Error
    }

    c.Content = content

    return Ok
}

func (c *Configure) GetContent() Content {
    if c.Content == nil {
        return nil
    }

    return c.Content
}
*/
/*
func (c *Configure) SetParser(parser Parser) int {
    if parser == nil {
       return Error
    }

    c.Parser = parser

    return c.Parser
}


func (c *Configure) GetParser() Parser {
    if c.Parser == nil {
        return nil
    }

    return c.Parser
}
*/
func (c *Configure) Materialized(cycle *Cycle) int {
    log := c.Log

    if configure := c.Get(); configure == nil {
        return Error
    }

    if c.value == nil {
        content := c.GetContent()
        if content == nil {
            log.Error("configure content: %s, filename: %s, size: %d\n",
                      content, c.GetFileName(), c.GetSize())

            return Error
        }

        if c.configure.Parser(content, &c.value) == Error {
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
        fmt.Println("unknown")
    }

    return Ok
}

func (c *Configure) doParse(materialized map[interface{}]interface{}, cycle *Cycle) int {
    log := c.Log

    flag := Ok

    modules := cycle.GetModules()
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
								    /*
            if module.Type != CONFIG_MODULE &&
               module.Type != c.moduleType {

                continue;
            }
            */

            commands := module.Commands;
            if commands == nil {
                continue;
            }

            //fmt.Printf("%s, %X, %X, %d\n", name, module.Type, c.moduleType, m)

            for i := 0; commands[i].Name.Len != 0; i++ {

                command := commands[i]

                if len(name) == command.Name.Len &&
                        name == command.Name.Data.(string) {

                				found = true;

                    if command.Type & c.commandType == 0 &&
                       command.Type & MAIN_CONFIG == 0 {

                        //flag = Error
																				    found = false
                        break
                    }

                    //fmt.Printf("h:%s, %X, %X\n", name, command.Type, c.commandType)

                    //log.Error("directive \"%s\" is not allowed here", name)
                    //					flag = Error
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

        if !found {
            log.Error("unkown")

            flag = Error
            break
        }

        if flag == Error {
            break
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

func (c *Configure) Parse(in []byte, out interface{}) int {
    fmt.Println("configure parser")
    return Ok
}
