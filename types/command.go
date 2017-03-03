/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

const (
    CONFIG_ARGS_NUMBER = 0x000000ff
    CONFIG_FLAG        = 0x00000200
    CONFIG_ANY         = 0x00000400
    CONFIG_1MORE       = 0x00000800
    CONFIG_2MORE       = 0x00001000
    CONFIG_BLOCK       = 0x00000100

    MAIN_CONFIG  =   0x01000000
    USER_CONFIG  =   0x0F000000
)

type SetFunc func(configure *AbstractConfigure, command *Command, cycle *AbstractCycle, config *unsafe.Pointer) string

type Command struct {
    Name    String
    Type    int
    Set     SetFunc
    Conf    int
    Offset  uintptr
    Post    interface{}
}

var NilCommand = Command{ NilString, 0, nil, 0, 0, nil }
