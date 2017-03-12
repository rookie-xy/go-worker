/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type Input struct {
    *Cycle
    *File
     input  InputIf
}

type InputIf interface {}

func NewInput() *Input {
    return &Input{}
}
