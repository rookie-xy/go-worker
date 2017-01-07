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
    return &yamlConfigure{ NewConfigure() }
}

type yamlConfigureContext struct {
    *AbstractContext
}

var Yaml = String{ len("yaml"), "yaml" }

var YamlConfigureContext = AbstractContext{
    Yaml,
    NewYamlConfigureContext().Context,
}

func NewYamlConfigureContext() *yamlConfigureContext {
    this := NewContext()

    if this == nil {
        return nil
    }

    this.Context = &yamlConfigureContext{}

    return &yamlConfigureContext{
        AbstractContext: this,
    }
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
