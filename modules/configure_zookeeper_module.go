/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
      "fmt"
    . "go-worker/types"
)

type zookeeperConfigure struct {
    *AbstractFile
}

func NewZookeeperConfigure() *zookeeperConfigure {
    return &zookeeperConfigure{}
}

func (zc *zookeeperConfigure) Open(name string) int {
    fmt.Println("from zookeeper configure")
    return Ok
}

func (zc *zookeeperConfigure) Read() int {
    fmt.Println("zookeeper configure read token")
    return Ok
}

type zookeeperConfigureContext struct {
    *AbstractContext
}

var Zookeeper = String{ len("zookeeper"), "zookeeper" }
var ZookeeperContext = NewZookeeperConfigureContext()

var ZookeeperConfigureContext = AbstractContext{
    Zookeeper,
    ZookeeperContext.Get(),
}

func NewZookeeperConfigureContext() *zookeeperConfigureContext {
    return &zookeeperConfigureContext{}
}

func (zcc *zookeeperConfigureContext) Get() Context {
    this := NewContext()
    if this == nil {
        return nil
    }

    this.Context = zcc

    return zcc.Set(this)
}

func (zcc *zookeeperConfigureContext) Set(context *AbstractContext) *zookeeperConfigureContext {
    if context == nil {
        return nil
    }

    zcc.AbstractContext = context

    return zcc
}

func (zcc *zookeeperConfigureContext) Create(cycle *AbstractCycle) unsafe.Pointer {
    configure := cycle.GetConfigure()
    if configure == nil {
        return nil
    }

    fileType := configure.GetFileType()
    if fileType == "" {
        return nil
    }

    if fileType != Zookeeper.Data.(string) {
       return nil
    }

    zc := NewZookeeperConfigure()
    if zc == nil {
        return nil
    }

    if configure.SetFile(zc) == Error {
        return nil
    }

    return unsafe.Pointer(zc)
}

func (zcc *zookeeperConfigureContext) Init(cycle *AbstractCycle, configure *unsafe.Pointer) string {
    fmt.Println("zookeeper configure init")
    return "0"
}

var ZookeeperConfigureModule = Module{
    0,
    0,
    &ZookeeperConfigureContext,
    nil,
    SYSTEM_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &ZookeeperConfigureModule)
}
