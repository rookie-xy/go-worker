/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

var (
    CONF_ARGS_NUMBER = 0x000000ff
    CONF_BLOCK  =     0x00000100
    CONF_FLAG   =     0x00000200
    CONF_ANY    =     0x00000400
    CONF_1MORE  =     0x00000800
    CONF_2MORE  =     0x00001000

    DIRECT_CONF  =    0x00010000

    MAIN_CONF    =    0x01000000
    ANY_CONF     =    0x0F000000
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
