/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractSignal struct {
    *AbstractFile
    signal  Signal
}

type Signal interface {
    Get() int
}
