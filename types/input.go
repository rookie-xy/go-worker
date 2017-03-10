/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type AbstractInput struct {
    *AbstractCycle
    *AbstractFile
     input  Input
}

type Input interface {}

func NewInput() *AbstractInput {
    return &AbstractInput{}
}
