/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type AbstractOutput struct {
    *AbstractCycle
    *AbstractFile
     output  Output
}

type Output interface {}

func NewOutput() *AbstractOutput {
    return &AbstractOutput{}
}
