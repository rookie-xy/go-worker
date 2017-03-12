/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
)

type ContextCreateFunc func(cycle *Cycle) unsafe.Pointer
type ContextInitFunc func(cycle *Cycle, configure *unsafe.Pointer) string

type Context struct {
    Name    String
    Create  ContextCreateFunc
    Init    ContextInitFunc
}
