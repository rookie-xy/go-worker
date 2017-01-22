/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractChannal struct {
    *AbstractRoutine
    name string
}

type Channal interface {
    Parse()
}
