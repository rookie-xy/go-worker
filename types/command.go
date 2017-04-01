/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

const (
    CONFIG_ARGS_NUMBER = 0x06000000
    CONFIG_FLAG        = 0x05000000
    CONFIG_ANY         = 0x04000000
    CONFIG_1MORE       = 0x03000000
    CONFIG_2MORE       = 0x02000000
    CONFIG_BLOCK       = 0x01000000
    CONFIG_VALUE       = 0x00000001

    MAIN_CONFIG  =   0x10000000
    USER_CONFIG  =   0xF0000000
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
