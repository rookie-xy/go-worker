/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractOutput struct {
    *AbstractRoutine
    name string
}

type Output interface {
    Parse()
}
