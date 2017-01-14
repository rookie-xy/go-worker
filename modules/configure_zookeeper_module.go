/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
    . "go-worker/types"
)

// 覆盖文件的读写方式
type zookeeperConfigure struct {
    *AbstractFile
}

func NewZookeeperConfigure() *zookeeperConfigure {
    return &zookeeperConfigure{}
}
