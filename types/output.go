/*
 * Copyright (C) 2017 Meng Shi
 */

package types

import (
    "unsafe"
    "strings"
)

type AbstractOutput struct {
    *AbstractCycle

     name   string
     output  Output
}

type Output interface {
    Parse()
}

func NewOutput() *AbstractOutput {
    return &AbstractOutput{}
}

type outputContext struct {
    *AbstractContext
}

var Stdout = String{ len("stdout"), "stdout" }
var OutputContext = NewOutputContext()

var OutputStdoutContext = AbstractContext{
    Stdout,
    OutputContext.Get(),
}

func NewOutputContext() *outputContext {
    return &outputContext{}
}

func (oc *outputContext) Get() Context {
    this := NewContext()
    if this == nil {
        return nil
    }

    this.Context = oc

    return oc.Set(this)
}

func (oc *outputContext) Set(context *AbstractContext) *outputContext {
        if context == nil {
                return nil
        }

        oc.AbstractContext = context

        return oc
}

func (oc *outputContext) Create(cycle *AbstractCycle) unsafe.Pointer {
    log := cycle.GetLog()

    configure := cycle.GetConfigure()
    if configure == nil {
        log.Info("aaa")
        return nil
    }

    fileName := configure.GetFileName()
    if fileName == "" {
        return nil
    }

    if !strings.HasSuffix(fileName, Stdin.Data.(string)) {
        return nil
    }

    o := NewOutput()
    if o == nil {
        return nil
    }

    /*
    if configure.Set(yc) == Error {
        return nil
    }
    */

    return unsafe.Pointer(o)
}

func (oc *outputContext) Init(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    //fmt.Println("input init for stdin")
    return "0"
}

var OutputModule = Module{
    0,
    0,
    &OutputStdoutContext,
    nil,
    USER_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &OutputModule)
}
