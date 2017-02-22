/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "fmt"
    "runtime"
)

type AbstractLog struct {
    *AbstractFile
     level  int
     log    Log
}

type Log interface {
    Debug(format string, d ...interface{})
    Info(format string, i ...interface{})
    Warn(format string, w ...interface{})
    Error(format string, e ...interface{})
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

func (l *AbstractLog) Debug(format string, d ...interface{}) {
    fmt.Printf(format, d)
    return
}

func (l *AbstractLog) Info(format string, i ...interface{}) {
    file := l.AbstractFile.GetFile()
    fmt.Fprintf(file, format, i)

    // TODO
    if file.Sync() != nil {
        //
    }

    return
}

func (l *AbstractLog) Warn(format string, w ...interface{}) {
    funcName, file, line, _ := runtime.Caller(0)
    fmts := format + runtime.FuncForPC(funcName).Name() + file + string(line)
    fmt.Printf(fmts, w)
    return
}

func (l *AbstractLog) Error(format string, e ...interface{}) {
    fmt.Printf(format, e)
    return
}
