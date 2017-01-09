/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
    "fmt"
)

type AbstractContext struct {
    Name     String
    Context  Context
}

type Context interface {
    Create(cycle *AbstractCycle) unsafe.Pointer
    Init(cycle *AbstractCycle, configure *unsafe.Pointer) string
}

func NewContext() *AbstractContext {
    return &AbstractContext{
        Context: &AbstractContext{},
    }
}

func (ac *AbstractContext) Create(cycle *AbstractCycle) unsafe.Pointer {
	fmt.Println("bbbbbbbbbbbbbbbbbb")
    return nil
}

func (ac *AbstractContext) Init(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    return ""
}
