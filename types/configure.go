/*
 * Copyright (C) 2016 Meng Shi
 */

package types

const (
    ConfNoargs = 0x00000001
    ConfTake1 = 0x00000002
    ConfTake2 = 0x00000004
    ConfTake3 = 0x00000008
    ConfTake4 = 0x00000010
    ConfTake5 = 0x00000020
    ConfTake6 = 0x00000040
    ConfTake7 = 0x00000080
    ConfTake12 = (ConfTake1|ConfTake2)

    ConfMore1 = 0x00000800
    ConfMore2 = 0x00001000
    ConfAny = 0x00000400

    ConfBlock = 0x00000100
    ConfFlag = 0x00000200

    MainConf = 0x01000000
    AnyConf = 0x0F000000
    DirectConf = 0x00010000
)

type AbstractConfigure struct {
    file  string
    name  string
}

type Configure interface {
    Set()
    Get()
    Parse() int
}

func NewConfigure() *AbstractConfigure {
    return &AbstractConfigure{}
}

func (c *AbstractConfigure) SetFile(file string) int {
    if file == "" {
        return Error
    }

    c.file = file

    return Ok
}

func (c *AbstractConfigure) GetFile() string {
    return c.file
}

func (c *AbstractConfigure) SetName(name string) int {
    if name == "" {
        return Error
    }

    c.name = name

    return Ok
}

func (c *AbstractConfigure) GetName() string {
    return c.name
}
