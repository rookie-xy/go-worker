/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type IO interface {
    Open(name string) int
    Close() int
    Read() int
    Write() int
    Type() *AbstractFile
}

func a() {
    close(aa);
}