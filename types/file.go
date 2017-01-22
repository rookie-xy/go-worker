/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "fmt"
    "os"
)

type AbstractFile struct {
    *AbstractLog
     file     *os.File
     name      string
     size      int64
     content   []byte
     action    Action
}

type Action interface {
    Open(name string) int
    Close() int
    Read() int
    Write() int
}

func NewFile(log *AbstractLog) *AbstractFile {
    return &AbstractFile{
        AbstractLog : log,
        file        : os.Stdout,
    }
}

func (f *AbstractFile) SetName(name string) int {
    if name == "" {
        return Error
    }

    f.name = name

    return Ok
}

func (f *AbstractFile) GetName() string {
    return f.name
}

func (f *AbstractFile) SetSize(size int64) int {
    if size < 0 {
        return Error
    }

    f.size = size

    return Ok
}

func (f *AbstractFile) GetSize() int64 {
    return f.size
}

func (f *AbstractFile) SetContent(content []byte) int {
    if content == nil {
        return Error
    }

    f.content = content

    return Ok
}

func (f *AbstractFile) GetContent() []byte {
    return f.content
}

func (f *AbstractFile) Set(action Action) int {
    if action == nil {
        return Error
    }

    f.action = action

    return Ok
}

func (f *AbstractFile) Get() Action {
    return f.action
}

func (f *AbstractFile) Open(name string) int {
    log := f.AbstractLog.Get()

    file, error := os.OpenFile(name, os.O_RDWR, 0666)
    if error != nil {
        log.Info("open file error: %s", error)
        return Error
    }

    f.file = file

    log.Info("from file open")

    return Ok
}

func (f *AbstractFile) Close() int {
    log := f.AbstractLog.Get()

    if error := f.file.Close(); error != nil {
        log.Info("close file error: %s", error)
        return Error
    }

    return Ok
}

func (f *AbstractFile) Read() int {
    var b []byte

    // TODO bugfix
    n, error := f.file.Read(b)
    if error != nil {
        return Error
    }
// TODO bugfix
    f.size = n
    f.content = b

    fmt.Println("from file read")

    return Ok
}

func (f *AbstractFile) Write() int {
    return Ok
}
