/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "fmt"

type AbstractLog struct {
    *AbstractFile
     log  Log
}

type Log interface {
    Debug()
    Info()
    Warn()
    Error()
}

func NewLog() *AbstractLog {
    return &AbstractLog{
        AbstractFile:NewFile(),
    }
}

func (l *AbstractLog) SetLog(log Log) int {
    if log == nil {
        return Error
    }

    l.log = log

    return Ok
}

func (l *AbstractLog) GetLog() Log {
    return l.log
}

func (l *AbstractLog) Debug() {
    fmt.Println("this is debug")
    return
}

func (l *AbstractLog) Info() {
    fmt.Println("this is info")
    return
}

func (l *AbstractLog) Warn() {
    fmt.Println("this is warn")
    return
}

func (l *AbstractLog) Error(error interface{}) {
    fmt.Println("this is error")
    return
}
