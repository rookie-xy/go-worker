/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractInput struct {
    *AbstractCycle
     name  string
}

type Input interface {
    Parse()
}
