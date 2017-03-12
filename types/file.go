/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "os"
)

type File struct {
    *Log
     file     *os.File
     name      string
     size      int64
     content   []byte
     action    IO
}

/*
type Action interface {
    Open(name string) int
    Close() int
    Read() int
    Write() int
    Type() *File
}
*/

func NewFile(log *Log) *File {
    return &File{
        Log : log,
        file        : os.Stdout,
    }
}

func (f *File) SetName(name string) int {
    if name == "" {
        return Error
    }

    f.name = name

    return Ok
}

func (f *File) GetName() string {
    return f.name
}

func (f *File) SetSize(size int64) int {
    if size < 0 {
        return Error
    }

    f.size = size

    return Ok
}

func (f *File) GetSize() int64 {
    return f.size
}

func (f *File) SetContent(content []byte) int {
    if content == nil {
        return Error
    }

    f.content = content

    return Ok
}

func (f *File) GetContent() []byte {
    return f.content
}

func (f *File) SetFile(file *os.File) int {
    if file == nil {
        return Error
    }

    f.file = file

    return Ok
}

func (f *File) GetFile() *os.File {
    return f.file
}

func (f *File) Set(action IO) int {
    if action == nil {
        return Error
    }

    f.action = action

    return Ok
}

func (f *File) Get() IO {
    return f.action
}

func (f *File) Open(name string) int {
    log := f.Log.Get()

    file, error := os.OpenFile(name, os.O_RDWR, 0777)
    if error != nil {
        log.Info("open file error: %s", error)
        return Error
    }

    stat, error := file.Stat()
    if error != nil {
        log.Info("stat file error: %s", error)
        return Error
    }

    f.file = file
    f.size = stat.Size()

    return Ok
}

func (f *File) Close() int {
    log := f.Log.Get()

    if error := f.file.Close(); error != nil {
        log.Info("close file error: %s", error)
        return Error
    }

    return Ok
}

func (f *File) Read() int {
    log := f.Log.Get()

    var char []byte

    if size := f.size; size <= 0 {
        log.Error("file size is: %d\n", size)
        return Error
    } else {
        char = make([]byte, size)
    }

    _, error := f.file.Read(char)
    if error != nil {
        log.Error("file read error: %s", error)
        return Error
    }

    f.content = char

    return Ok
}

func (f *File) Write() int {
    return Ok
}

func (f *File) Type() *File {
    return f
}
