/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type Input struct {
    *AbstractFile

     input  InputIf
}

type InputIf interface {}

func NewInput() *Input {
    return &Input{}
}

func (i *Input) SetInput() {

}
