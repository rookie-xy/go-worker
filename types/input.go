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

var (
    User = String{ len("user"), "user" }
    InputContext = NewInputContext()
)

var InputStdinContext = AbstractContext{
    User,
    InputContext.Get(),
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
    Stdin = String{ len("stdin"), "stdin" }
    input AbstractInput
)

var InputCommands = []Command{

    { Stdin,
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

var InputModule = Module{
    0,
    0,
    &InputStdinContext,
    InputCommands,
    USER_MODULE,
    initInputModule,
    routineInputModule,
}

func initInputModule(cycle *AbstractCycle) int {
    return Ok
}

func routineInputModule(cycle *AbstractCycle) int {

    for ;; {
        fmt.Println("aaaaaaaaaaa")
    }

    return Ok
}

func init() {
    Modules = append(Modules, &InputModule)
}
