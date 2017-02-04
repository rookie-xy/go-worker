/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractChannal struct {
    *AbstractCycle
    name string
}

type Channal interface {
    Parse()
}
