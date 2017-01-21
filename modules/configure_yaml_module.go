/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "fmt"
      "unsafe"
      "strings"
      "log"
      "gopkg.in/yaml.v2"

    . "go-worker/types"
)

type yamlConfigure struct {
    *AbstractConfigure
}

func NewYamlConfigure() *yamlConfigure {
    return &yamlConfigure{}
}

func (yc *yamlConfigure) Parse() int {
    var data = `
a: Easy!
b:
c: 2
d: [3, 4]`

    t := make(map[string]interface{})

    err := yaml.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    a := t["d"].([]interface{})

    fmt.Printf("yaml configure parse: %d\n", a[1])
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

    yc := NewYamlConfigure()
    if yc == nil {
        return nil
    }

    if configure.Set(yc) == Error {
        return nil
    }

    return unsafe.Pointer(yc)
}

func (ycc *yamlConfigureContext) Init(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    fmt.Println("yaml configure init")
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
