/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "fmt"
    "runtime"
)

type Log struct {
    *File
     level  int
     log    LogIf
}

type LogIf interface {
    Debug(format string, d ...interface{})
    Info(format string, i ...interface{})
    Warn(format string, w ...interface{})
    Error(format string, e ...interface{})
}

var Level = [...]string{ "stderr", "emerg", "alert", "crit", "error",
                         "warn", "notice", "info", "debug" }

func NewLog() *Log {
    return &Log{
        File : NewFile(nil),
        level        : 4,
    }
}

func (l *Log) Set(log LogIf) int {
    if log == nil {
        return Error
    }

    l.log = log

    return Ok
}

func (l *Log) Get() LogIf {
    return l.log
}

func (l *Log) Debug(format string, d ...interface{}) {
    fmt.Printf(format, d)
    return
}

func (l *Log) Info(format string, i ...interface{}) {
    file := l.File.GetFile()
    fmt.Fprintf(file, format, i)

    // TODO
    if file.Sync() != nil {
        //
    }

    return
}

func (l *Log) Warn(format string, w ...interface{}) {
    funcName, file, line, _ := runtime.Caller(0)
    fmts := format + runtime.FuncForPC(funcName).Name() + file + string(line)
    fmt.Printf(fmts, w)
    return
}

func (l *Log) Error(format string, e ...interface{}) {
    fmt.Printf(format, e)
    return
}
