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
    Stop() int

    Monitor() int
}

func (f MainFunc) Start() int {
    go f(nil)
    return Ok
}

func (f MainFunc) Stop() int {
    return Ok
}

func (r *AbstractRoutine) Monitor() int {
    return Ok
}
