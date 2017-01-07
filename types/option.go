/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type AbstractOption struct {
    argc   int
    argv   []string
    items  map[string]interface{}
}

type Option interface {
    /*
     * This is option interface, it is
     * to be any impl
     */
    Parse() int
}

func NewOption() *AbstractOption {
    return &AbstractOption{
        items: make(map[string]interface{}),
    }
}

func (ao *AbstractOption) GetArgc() int {
    return ao.argc
}

func (ao *AbstractOption) GetArgv() []string {
    return ao.argv
}

func (ao *AbstractOption) SetArgs(argc int, argv []string) int {
    if argc <= 0 || argv == nil {
        return Error
    }

    ao.argc = argc
    ao.argv = argv

    return Ok
}

func (ao *AbstractOption) SetItem(k string, v interface{}) {
    ao.items[k] = v
}

func (ao *AbstractOption) GetItem(k string) interface{} {
    return ao.items[k]
}

func (ao *AbstractOption) Parse() int {
    return Ok
}
