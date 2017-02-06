/*
 * Copyright (C) 2017 Meng Shi
 */

package types

import (
    "unsafe"
    "fmt"
)

type AbstractInput struct {
    *AbstractCycle
    *AbstractFile
     stdin  bool
     input  Input
}

type Input interface {
    Parse()
}

func NewInput() *AbstractInput {
    return &AbstractInput{}
}

type inputContext struct {
    *AbstractContext
}

var user = String{ len("user"), "user" }

var inputStdinContext = &AbstractContext{
    user,
    NewInputContext().Get(),
}

func NewInputContext() *inputContext {
    return &inputContext{}
}

func (ic *inputContext) Get() Context {
    this := NewContext()
    if this == nil {
        return nil
    }

    this.Context = ic

    return ic.Set(this)
}

func (ic *inputContext) Set(context *AbstractContext) *inputContext {
    if context == nil {
        return nil
    }

    ic.AbstractContext = context

    return ic
}

func (ic *inputContext) Create(cycle *AbstractCycle) unsafe.Pointer {
    log := cycle.GetLog()

    i := NewInput()
    if i == nil {
	log.Error("new input error")
        return nil
    }

    i.stdin = false
    i.AbstractCycle = cycle

    return unsafe.Pointer(i)
}

func (i *inputContext) Init(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    log := cycle.GetLog()

    input := (*AbstractInput)(unsafe.Pointer(configure))
    if input.stdin == true {
        input.AbstractFile = NewFile(log)
    }

    return "0"
}

var (
    stdin = String{ len("stdin"), "stdin" }
    input AbstractInput
)

var inputCommands = []Command{

    { stdin,
      0,
      configureSetFlag,
      0,
      unsafe.Offsetof(input.stdin),
      nil },

    NilCommand,
}

func configureSetFlag(cf *AbstractConfigure, command *Command, conf interface{}) string {
    return ""
}

var inputModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    inputStdinContext,
    inputCommands,
    USER_MODULE,
    inputInit,
    inputMain,
}

func inputInit(cycle *AbstractCycle) int {
    return Ok
}

func inputMain(cycle *AbstractCycle) int {

    for ;; {
        fmt.Println("aaaaaaaaaaa")
    }

    return Ok
}

func init() {
    Modules = append(Modules, &inputModule)
}
