/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
    "fmt"
    "sync"
    "math/rand"
    "strings"
    "strconv"
)

type RoutineFunc func(p unsafe.Pointer) int

type Routine struct {
    *Log
     sync.Mutex

     name     string
     number   int
     events   map[string]*Event
}

type RoutineIf interface {
    Start(cycle *Cycle) int
    Stop() int

    Monitor() int
}

func NewRoutine() *Routine {
    return &Routine{}
}

func (r *Routine) Go(name string, fn interface{}, args ...interface{}) {
    go func(this *Routine, n string, fn interface{}, args ...interface{} ) {
        if this.register(name) == Error {
            return Error
        }

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

        /*
        if len(args) > 1 {
            fn.(func(...interface{}) int)(args)
        } else if len(args) == 1 {
            fn.(func(interface{}) int)(args[0])
        } else {
            fn.(func() int)()
        }
        */

    }(r, name, fn, args...)
}

func (r *Routine) register(name string) int {
    event := NewEvent()

    event.name = name
    event.gid = uint64(rand.Int63())
    event.notice = make(chan *Event)

    r.Lock()
    defer r.Unlock()

    if r.events == nil {
        r.events = make(map[string]*Event)

    } else if _, ok := r.events[event.name]; ok {
        r.Warn("goroutine channel already defined: %q", event.name)
        return Ignore
    }

    r.events[event.name] = event

    return Ok
}

func (r *Routine) unregister(name string) int {
    r.Lock()
    defer r.Unlock()

    if _, ok := r.events[name]; !ok {
        r.Warn("goroutine channel not find: %q", name)
        return Error
    }

    delete(r.events, name)

    return Ok
}

func (r *Routine) Start(cycle *Cycle) int {
    return Ok
}

func (r *Routine) Stop() int {
    return Ok
}

func (r *Routine) Monitor(name string) {

    select {

    case notice := <-r.events[name].notice:

        gid := notice.gid
        op  := notice.opcode

        if gid == r.events[name].gid {

            switch op {

            case KILL:
                r.unregister(name)
                r.Warn("gid[" + gid + "] quit")
                return

            default:
                r.Info("unknown signal")

            }
        }

    default:
        fmt.Println("no signal")
    }
}
