/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type AbstractChannal struct {
    *AbstractCycle
    *AbstractFile
     channal  Channal
}

type Channal interface {
    push()
    pull()
}

func NewChannal() *AbstractChannal {
    return &AbstractChannal{}
}
