/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
    . "go-worker/types"
    "fmt"
    "unsafe"
)

type yamlConfigure struct {
    *AbstractConfigure
}

func NewYamlConfigure() *yamlConfigure {
    return &yamlConfigure{}
}

func (yc *yamlConfigure) Parse() int {
    fmt.Println("yaml configure parse")
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
    YamlContext.Override(),
}

func NewYamlConfigureContext() *yamlConfigureContext {
    return &yamlConfigureContext{}
}

func (ycc *yamlConfigureContext) Override() Context {
    this := NewContext()
    if this == nil {
        return nil
    }

    this.Context = ycc

    return ycc.Set(this)
}

func (ycc *yamlConfigureContext) Set(context *AbstractContext) *yamlConfigureContext {
    ycc.AbstractContext = context
    return ycc
}

func (ycc *yamlConfigureContext) Create(cycle *AbstractCycle) unsafe.Pointer {
    configure := cycle.GetConfigure()
    if configure == nil {
        return nil
    }

    if configure.GetTypeName() == Yaml.Data {
        fmt.Println("is right")
    }

    fmt.Println("sssssssssssss")

    yc := NewYamlConfigure()
    if yc == nil {
        return nil
    }

    fmt.Println("qqqqqqqqqqqqqqq")
    if configure.Set(yc) == Error {
        return nil
    }
/*
    if yc.Override() == Error {
        return nil
    }
    */

    fmt.Println("ooooooooooooooooooooo")

    return nil
}

func (ycc *yamlConfigureContext) Init(cycle *AbstractCycle) string {
    return ""
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
