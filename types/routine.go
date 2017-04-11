/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

type RoutineFunc func(p unsafe.Pointer) int

type Routine struct {
    name    string
    number  int
}

type RoutineIf interface {
    Start(cycle *Cycle) int
    Stop() int

    Monitor() int
}

func NewRoutine() *Routine {
    return &Routine{}
}

func (r *Routine) Go(flag int, rf RoutineFunc, p unsafe.Pointer) int {
    go rf(p)

    return Ok
}

func (r *Routine) Start(cycle *Cycle) int {
    return Ok
}

func (r *Routine) Stop() int {
    return Ok
}

func (r *Routine) Monitor() int {
    return Ok
}
