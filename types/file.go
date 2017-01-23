/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "os"
)

type AbstractFile struct {
    *AbstractLog
     file     *os.File
     name      string
     size      int
     content   []byte
     action    Action
}

type Action interface {
    Open(name string) int
    Close() int
    Read() int
    Write() int
    Type() *AbstractFile
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

func (f *AbstractFile) SetSize(size int) int {
    if size < 0 {
        return Error
    }

    f.size = size

    return Ok
}

func (f *AbstractFile) GetSize() int {
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

    file, error := os.OpenFile(name, os.O_RDWR, 0777)
    if error != nil {
        log.Info("open file error: %s", error)
        return Error
    }

    f.file = file

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
    log := f.AbstractLog.Get()
    var b []byte = make([]byte, 1024)

    n, error := f.file.Read(b)
    if error != nil {
        log.Error("file read error: %s", error)
        return Error
    }

    f.size = n
    f.content = b

    return Ok
}

func (f *AbstractFile) Write() int {
    return Ok
}

func (f *AbstractFile) Type() *AbstractFile {
    return f
}
