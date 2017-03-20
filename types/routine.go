/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type Routine struct {
    name    string
    number  int
}

type RoutineIf interface {
    Start(cycle *Cycle) int
    Stop() int

    Monitor() int
}

func (f MainFunc) Start(cycle *Cycle) int {
    if f == nil {
        return Error
    }

    go f(cycle)

    return Ok
}

func (f MainFunc) Stop() int {
    return Ok
}

func (r *Routine) Monitor() int {
    return Ok
}
