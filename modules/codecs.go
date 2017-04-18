/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    CODEC_MODULE = 0x70000000
    CODEC_CONFIG = MAIN_CONFIG|CONFIG_BLOCK
)

var codec = String{ len("codec"), "codec" }
var codecContext = &Context{
    codec,
    nil,
    nil,
}

var codecs = String{ len("codecs"), "codecs" }
var codecCommands = []Command{

    { codecs,
      CODEC_CONFIG,
      codecBlock,
      0,
      0,
      nil },

    NilCommand,
}

func codecBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    if cycle == nil {
        return Error
    }

    flag := CODEC_CONFIG|CONFIG_MAP
    cycle.Block(cycle, CODEC_MODULE, flag)

    return Ok
}

var codecModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(codecContext),
    codecCommands,
    CONFIG_MODULE,
    nil,
    nil,
    nil,
}

func init() {
    Modules = Load(Modules, &codecModule)
}