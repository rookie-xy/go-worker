/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
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
    flag := Ok

    go func(this *Routine, n string, fn interface{}, args ...interface{} ) {

        switch len(args) {

        case 0:
            flag = fn.(func() int)()

        case 1:
            flag = fn.(func(interface{}) int)(args[0])

        default:
            flag = fn.(func(...interface{}) int)(args)
        }

    }(r, name, fn, args...)

    if flag == Error {
        r.Warn(name)
        return -1
    }

    gid, rv := r.Register(name)
    if rv != Ok  {
        return -1
    }

    return gid
}

func (r *Routine) Register(name string) (int64, int) {
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

func (r *Routine) Unregister(id int64) int {
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

func (r *Routine) Check(id int64, flag uint8) int {
    if id < 0 {
        r.Warn("The id error")
    }

    switch flag {

    case BLOCKING:
        //<-r.events[name].notice

    case NONBLOCKING:
        /*
        select {

        case 1://<-r.events[name].notice:

        default:
            return -1
        }
        */
        return -1
    }

    return 1
}

func (r *Routine) GetEvent(id int64) *Event {
    if e, ok := r.events[id]; ok {
        return e
    }

    return nil
}
