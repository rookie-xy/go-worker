/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
)

type ContextCreateFunc func(cycle *AbstractCycle) unsafe.Pointer
type ContextInitFunc func(cycle *AbstractCycle, configure *unsafe.Pointer) string

type AbstractContext struct {
    Name    String
    Create  ContextCreateFunc
    Init    ContextInitFunc
}
