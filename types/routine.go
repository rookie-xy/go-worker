/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractRoutine struct {
    name    string
    number  int
}

type Routine interface {
    Manage() int
    Status() int
}
