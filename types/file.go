/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "fmt"

type AbstractFile struct {
    name      string
    size      int64
    content   []byte
    file      File
}

type File interface {
    Open() int
    Close() int
    Read() int
    Write() int
}

func NewFile() *AbstractFile {
    return &AbstractFile{}
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

func (f *AbstractFile) SetFile(file File) int {
    if file == nil {
        return Error
    }

    f.file = file

    return Ok
}

func (f *AbstractFile) GetFile() File {
    return f.file
}

func (f *AbstractFile) Open() int {
    fmt.Println("from file open")
    return Ok
}

func (f *AbstractFile) Close() int {
    return Ok
}

func (f *AbstractFile) Read() int {
	fmt.Println("from file read")
    return Ok
}

func (f *AbstractFile) Write() int {
    return Ok
}
