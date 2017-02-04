/*
 * Copyright (C) 2017 Meng Shi
 */

package types

import (
    "unsafe"
    "strings"
)

type AbstractChannal struct {
    *AbstractCycle

     name    string
     channal  Channal
}

type Channal interface {
    Parse()
}

func NewChannal() *AbstractChannal {
    return &AbstractChannal{}
}

type channalContext struct {
    *AbstractContext
}

var Memory = String{ len("memory"), "memory" }
var ChannalContext = NewChannalContext()

var ChannalMemoryContext = AbstractContext{
    Memory,
    ChannalContext.Get(),
}

func NewChannalContext() *channalContext {
    return &channalContext{}
}

func (cc *channalContext) Get() Context {
    this := NewContext()
    if this == nil {
        return nil
    }

    this.Context = cc

    return cc.Set(this)
}

func (cc *channalContext) Set(context *AbstractContext) *channalContext {
    if context == nil {
        return nil
    }

    cc.AbstractContext = context

    return cc
}

func (cc *channalContext) Create(cycle *AbstractCycle) unsafe.Pointer {
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

    c := NewChannal()
    if c == nil {
        return nil
    }

    /*
    if configure.Set(yc) == Error {
        return nil
    }
    */

    return unsafe.Pointer(c)
}

func (cc *channalContext) Init(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    //fmt.Println("input init for stdin")
    return "0"
}

var ChannalModule = Module{
    0,
    0,
    &ChannalMemoryContext,
    nil,
    USER_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &ChannalModule)
}