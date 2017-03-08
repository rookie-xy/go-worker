/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "unsafe"
)

var (
    ConfigOk    =  0
    ConfigError = -1
)

type AbstractConfigure struct {
    *AbstractLog
    *AbstractFile

     resource     string
     fileName     string
     commandType  int
     moduleType   int64
     value        interface{}
     configure    Configure
}

type Configure interface {
    Parse() int
    ReadToken() int
}

func NewConfigure(log *AbstractLog) *AbstractConfigure {
    return &AbstractConfigure{
        AbstractLog  : log,
        AbstractFile : NewFile(log),
    }
}

func (c *AbstractConfigure) SetName(file string) int {
    if file == "" {
        return Error
    }

    if c.AbstractFile.SetName(file) == Error {
        return Error
    }

    return Ok
}

func (c *AbstractConfigure) GetName() string {
    return c.AbstractFile.GetName()
}

func (c *AbstractConfigure) SetFileName(fileName string) int {
    if fileName == "" {
        return Error
    }

    c.fileName = fileName

    return Ok
}

func (c *AbstractConfigure) GetFileName() string {
    return c.fileName
}

func (c *AbstractConfigure) SetFileType(fileType string) int {
    if fileType == "" {
        return Error
    }

    if c.SetName(fileType) == Error {
        return Error
    }

    return Ok
}

func (c *AbstractConfigure) GetFileType() string {
    if fileType := c.GetName(); fileType != "" {
        return fileType
    }

    return ""
}

func (c *AbstractConfigure) SetFile(action IO) int {
    if action == nil {
        return Error
    }

    if c.AbstractFile.Set(action) == Error {
        return Error
    }

    return Ok
}

func (c *AbstractConfigure) GetFile() IO {
    if file := c.AbstractFile.Get(); file != nil {
        return file
    }

    return nil
}

func (c *AbstractConfigure) SetResource(resource string) int {
    if resource == "" {
        return Error
    }

    c.resource = resource

    return Ok
}

func (c *AbstractConfigure) GetResource(resource string) string {
    return c.resource
}

func (c *AbstractConfigure) Get() Configure {
    log := c.AbstractLog.Get()


    file := c.AbstractFile.Get()
    if file == nil {
        file = NewFile(c.AbstractLog)
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

func (c *AbstractConfigure) Set(configre Configure) int {
    if configre == nil {
        return Error
    }

    c.configure = configre

    return Ok
}

func (c *AbstractConfigure) SetModuleType(moduleType int64) int {
    if moduleType <= 0 {
        return Error
    }

    c.moduleType = moduleType

    return Ok
}

func (c *AbstractConfigure) SetCommandType(commandType int) int {
    if commandType <= 0 {
        return Error
    }

    c.commandType = commandType

    return Ok
}

func (c *AbstractConfigure) GetValue() interface{} {
   return c.value
}

func SetFlag(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
    if config == nil {
        return "0"
    }

    pointer := (*bool)(unsafe.Pointer(uintptr(*config) + command.Offset))
    if pointer == nil {
        return "0"
    }

    flag := configure.GetValue()
    if flag == true {
        *pointer = true
    } else if flag == false {
        *pointer = false
    } else {
        return "-1"
    }

    /*
    if command.Post != nil {
        post := command.Post.(DvrConfPostType);
        post.Handler(cf, post, *p);
    }
    */

    return ""
}

func SetString(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {
    pointer := (*string)(unsafe.Pointer(uintptr(*config) + command.Offset))
    if pointer == nil {
        return "0"
    }

    strings := configure.GetValue()
    if strings == nil {
        return "0"
    }

//    fmt.Printf("configureSetString: %s\n", *pointer)
    *pointer = strings.(string)

    return "0"
}

func SetNumber(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string {

    pointer := (*int)(unsafe.Pointer(uintptr(*config) + command.Offset))
    if pointer == nil {
        return "0"
    }

    number := configure.GetValue()
    if number == nil {
        return "0"
    }

//    fmt.Printf("configureSetNumber: %d\n", *pointer)
    *pointer = number.(int)

    return "0"
}

func (c *AbstractConfigure) Parse(cycle *AbstractCycle) int {
    log := c.AbstractLog.Get()

    if configure := c.Get(); configure != nil {
        if configure.Parse() == Error {
            return Error
        }

        return Ok
    }

    // TODO default process
    if c.value == nil {
        content := c.GetContent()
        if content == nil {
            log.Error("configure content: %s, filename: %s, size: %d\n",
                      content, c.GetFileName(), c.GetSize())

            return Error
        }

        error := yaml.Unmarshal(content, &c.value)
        if error != nil {
            log.Error("yanm unmarshal error: %s\n", error)
            return Error
        }
    }

    switch v := c.value.(type) {

    case []interface{} :
        for _, value := range v {
            c.value = value
            c.Parse(cycle)
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

func (c *AbstractConfigure) doParse(materialized map[interface{}]interface{}, cycle *AbstractCycle) int {
    log := c.AbstractLog.Get()

    flag := Ok

    for key, value := range materialized {

        if key != nil && value != nil {
            flag = Ok
        }

        name := key.(string)
        found := false

        for m := 0; flag != Error && !found && Modules[m] != nil; m++ {
            module := Modules[m]
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

                    //log.Error("directive \"%s\" is not allowed here", name)
                    //					flag = Error
                    context := cycle.GetContext(module.Index)

                    c.value = value
                    command.Set(c, &command, cycle, context)
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

func (c *AbstractConfigure) ReadToken() int {
    fmt.Println("configure read token")
    return Ok
}
