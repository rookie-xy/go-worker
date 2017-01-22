/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "fmt"

type AbstractLog struct {
    *AbstractFile
     level  int
     log    Log
}

type Log interface {
    Debug(d ...interface{})
    Info(i ...interface{})
    Warn(w ...interface{})
    Error(e ...interface{})
}

var Level = [...]string{ "stderr", "emerg", "alert", "crit", "error",
            "warn", "notice", "info", "debug" }

func NewLog() *AbstractLog {
    return &AbstractLog{
        AbstractFile : NewFile(nil),
        level        : 4,
    }
}

func (l *AbstractLog) Set(log Log) int {
    if log == nil {
        return Error
    }

    l.log = log

    return Ok
}

func (l *AbstractLog) Get() Log {
    return l.log
}

func (l *AbstractLog) Debug(debug interface{}) {
    fmt.Println("this is debug")
    return
}

func (l *AbstractLog) Info(i ...interface{}) {
    fmt.Println("this is info")
    return
}

func (l *AbstractLog) Warn(warn interface{}) {
    fmt.Println("this is warn")
    return
}

func (l *AbstractLog) Error(error interface{}) {
    fmt.Println("this is error")
    return
}
