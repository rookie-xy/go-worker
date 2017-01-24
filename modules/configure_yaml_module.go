/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "fmt"
      "unsafe"
      "strings"
//      "log"
//      "gopkg.in/yaml.v2"

    . "go-worker/types"
    "gopkg.in/yaml.v2"
)

type yamlConfigure struct {
    *AbstractConfigure
    commandType  uint
    moduleType   uint
}

func NewYamlConfigure(configure *AbstractConfigure) *yamlConfigure {
    return &yamlConfigure{
        AbstractConfigure:configure,
    }
}

func (yc *yamlConfigure) Parse() int {
    log := yc.AbstractLog.Get()

    content := yc.AbstractConfigure.GetContent()
    if content == nil {
        log.Error("configure content: %s, filename: %s, size: %d\n",
                  content, yc.AbstractConfigure.GetFileName(),
                  yc.AbstractConfigure.GetSize())

        return Error
    }

    materialized := make(map[interface{}]interface{})

    error := yaml.Unmarshal(content, &materialized)
    if error != nil {
        log.Error("yanm unmarshal error: %s\n", error)
        return Error
    }

    for key, value := range materialized {
        found := false
        flag := Ok
        name := key.(string)

        for i := 0; flag != Error && !found && Modules[i] != nil; i++ {
            module := Modules[i]

            if module.Type != SYSTEM_MODULE &&
                              module.Type != yc.moduleType {
                continue;
            }

            commands := module.Commands;
            if commands == nil {
                continue;
            }

            for i = 0; commands[i].Name.Len != 0; i++ {
                command := commands[i]

                if len(name) == command.Name.Len &&
                        name == command.Name.Data.(string) {

                    found = true;

                    if command.Type & yc.commandType == 0 {
                        log.Error("directive \"%s\" is not allowed here",
                                   name)
                        flag = Error
                        break;
                    }

                    if value == nil {
                        log.Error("lllll: %d\n", 10)
                    }
/*
                    if cmd.Type & DvrConfAny != 0 {
                        valid = true;

                    } else if cmd.Type & DvrConfFlag != 0 {
                        if cf.Args.Nelts == 2 {
                            valid = true;
                        } else {
                            valid = false;
                        }

                    } else if cmd.Type & DvrConfMore1 != 0 {
                        if cf.Args.Nelts > 1 {
                            valid = true;
                        } else {
                            valid = false;
                        }

                    } else if cmd.Type & DvrConfMore2 != 0 {
                        if cf.Args.Nelts > 2 {
                            valid = true;
                        } else {
                            valid = false;
                        }

                    } else if cf.Args.Nelts <= 10 &&
                            (cmd.Type & DvrArgumentNumber[cf.Args.Nelts - 1] != 0) {
                        valid = true;

                    } else {
                        valid = false;
                    }

                    if !valid {
                        DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                            "invalid number arguments in directive \"%s\" in %s:%d",
                            name.Data,
                            cf.ConfFile.File.Name.Data,
                            cf.ConfFile.Line);
                        rc = DvrError;
                        break;
                    }

                    conf = nil;

                    if command.Type & DvrDirectConf != 0 {
                        conf = cf.Ctx[DvrModules[m].Index];

                    } else if command.Type & DvrMainConf != 0 {
                        conf = cf.Ctx[DvrModules[m].Index];

                    } else if cf.Ctx != nil {
                        conf = cf.Ctx[DvrModules[m].CtxIndex];
                    }

                    rv := command.Set(cf, &command, value);

                    if rv == DvrConfOk {
                        break;

                    } else if rv == DvrConfError {
                        rc = DvrError;
                        break;

                    } else {
                        DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                            "the \"%s\" directive %s in %s:%d",
                            name.Data, rv,
                            cf.ConfFile.File.Name.Data,
                            cf.ConfFile.Line);
                        rc = DvrError;
                        break;
                    }
                    */
                }
            }
        }
    }

    return Ok
}

func (yc *yamlConfigure) ReadToken() int {
    fmt.Println("yaml configure read token")
    return Ok
}

type yamlConfigureContext struct {
    *AbstractContext
}

var Yaml = String{ len("yaml"), "yaml" }
var YamlContext = NewYamlConfigureContext()

var YamlConfigureContext = AbstractContext{
    Yaml,
    YamlContext.Get(),
}

func NewYamlConfigureContext() *yamlConfigureContext {
    return &yamlConfigureContext{}
}

func (ycc *yamlConfigureContext) Get() Context {
    this := NewContext()
    if this == nil {
        return nil
    }

    this.Context = ycc

    return ycc.Set(this)
}

func (ycc *yamlConfigureContext) Set(context *AbstractContext) *yamlConfigureContext {
    if context == nil {
        return nil
    }

    ycc.AbstractContext = context

    return ycc
}

func (ycc *yamlConfigureContext) Create(cycle *AbstractCycle) unsafe.Pointer {
    configure := cycle.GetConfigure()
    if configure == nil {
        return nil
    }

    fileName := configure.GetFileName()
    if fileName == "" {
        return nil
    }

    if !strings.HasSuffix(fileName, Yaml.Data.(string)) {
        return nil
    }

    yc := NewYamlConfigure(configure)
    if yc == nil {
        return nil
    }

    if configure.Set(yc) == Error {
        return nil
    }

    return unsafe.Pointer(yc)
}

func (ycc *yamlConfigureContext) Init(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    //fmt.Println("yaml configure init")
    return "0"
}

var YamlConfigureModule = Module{
    0,
    0,
    &YamlConfigureContext,
    nil,
    SYSTEM_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &YamlConfigureModule)
}
