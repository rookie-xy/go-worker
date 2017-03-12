/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type Channal struct {
    *Cycle
    *File
     channal  ChannalIf
}

type ChannalIf interface {
    push()
    pull()
}

func NewChannal() *Channal {
    return &Channal{}
}
