/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractRoutine struct {
    name    string
    number  int
}

type Routine interface {
    Start() int
}

func (f MainFunc) Start() int {
    go f(nil)
    return Ok
}
