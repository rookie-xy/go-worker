/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

const (
    CONFIG_BLOCK = 0x00100000
    CONFIG_MAP   = 0x00200000
    CONFIG_ARRAY = 0x00300000
    CONFIG_VALUE = 0x00F00000

    MAIN_CONFIG  = 0x01000000
    USER_CONFIG  = 0x0F000000
)

type SetFunc func(cycle *Cycle, cmd *Command, p *unsafe.Pointer) int

type Command struct {
    Name    String
    Type    int
    Set     SetFunc
    Conf    int
    Offset  uintptr
    Post    interface{}
}

var NilCommand = Command{ NilString, 0, nil, 0, 0, nil }
