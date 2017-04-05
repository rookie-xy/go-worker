/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    CODEC_MODULE = 0x6F6300000000
    CODEC_CONFIG = 0x01000000
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
      MAIN_CONFIG|CONFIG_MAP,
      codecBlock,
      0,
      0,
      nil },

    NilCommand,
}

func codecBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    cycle.Configure.Block(CODEC_MODULE, CODEC_CONFIG)
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
}

func init() {
    //Modules = append(Modules, &codecModule)
}