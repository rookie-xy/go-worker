/*
 * Copyright (C) 2016 Meng Shi
 */

package types

var (
    Ok     =  0
    Error  = -1
    Again  = -2
    Ignore = -3
)

const (
    MODULE_V1  = 0
    CONTEXT_V1 = 0

    SYSTEM_MODULE = 0x10000000
    CONFIG_MODULE = 0xF0000000
)

const (
    CONFIG_BLOCK = 0x00100000
    CONFIG_MAP   = 0x00200000
    CONFIG_ARRAY = 0x00300000
    CONFIG_VALUE = 0x00F00000

    MAIN_CONFIG  = 0x01000000
    USER_CONFIG  = 0x0F000000
)

/* magic */
const (
    MESSAGE = 0x01
    NOTICE  = 0x02
    PACKAGE = 0X03
)

/* opcode */
const (
    START uint8 = iota
    STOP
    LOAD
    RELOAD
    RESTART
    KILL
)

/* status */
const (
    NONBLOCKING uint8 = iota
    BLOCKING
)
