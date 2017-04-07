/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type Output struct {
    *Cycle
    *AbstractFile

     output  OutputIf
}

type OutputIf interface {}

func NewOutput() *Output {
    return &Output{}
}
