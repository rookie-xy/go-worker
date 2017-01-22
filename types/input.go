/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractInput struct {
    *AbstractRoutine
    name  string
}

type Input interface {
    Parse()
}
