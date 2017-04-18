/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
    "fmt"
    "sync"
    "math/rand"
)

type RoutineFunc func(p unsafe.Pointer) int

type Routine struct {
    *Log

     sync.Mutex
     events   map[int64]*Event
}

type RoutineIf interface {
    Start(cycle *Cycle) int
    Stop() int

    Monitor() int
}

func NewRoutine() *Routine {
    return &Routine{
        events: make(map[int64]*Event),
    }
}

func (r *Routine) Go(name string, fn interface{}, args ...interface{}) int64 {
    gid, flag := r.register(name)
    if flag == Error {
        return gid
    }

    go func(this *Routine, n string, fn interface{}, args ...interface{} ) {
        switch len(args) {

        case 0:
            if fn.(func() int)() == Error {
               r.Warn("f")
            }

        case 1:
            if fn.(func(interface{}) int)(args[0]) == Error {
                r.Warn("f")
            }

        default:
            fn.(func(...interface{}) int)(args)
        }

    }(r, name, fn, args...)

    return gid
}

func (r *Routine) register(name string) (int64, int) {
    routine := NewEvent()

    routine.id = rand.Int63()
    routine.name = name
    routine.magic = NOTICE

    //routine.notice = make(chan *Event)

    r.Lock()
    defer r.Unlock()

    if _, ok := r.events[routine.id]; ok {
        r.Warn("goroutine channel already defined: %s, %d",
                name, routine.id)
        return routine.id, Ignore
    }

    r.events[routine.id] = routine

    return routine.id, Ok
}

func (r *Routine) unregister(id int64) int {
    r.Lock()
    defer r.Unlock()

    if _, ok := r.events[id]; !ok {
        r.Warn("goroutine id not find: %d", id)
        return Error
    }

    delete(r.events, id)

    return Ok
}

func (r *Routine) Start(cycle *Cycle) int {
    return Ok
}

func (r *Routine) Stop() int {
    return Ok
}

func (r *Routine) Check(name string) {
    if name == "" {
        r.Warn("The name not found, is null")
    }

    select {
/*
    case notice := <-r.events[name].notice:

        gid := notice.gid
        op  := notice.opcode

        if gid == r.events[name].gid {

            switch op {

            case KILL:
                r.unregister(name)
                r.Warn("gid is :%d\n", gid)
                return

            default:
                r.Info("unknown signal")
            }
        }
        */

    default:
        fmt.Println("no signal")
    }
}
