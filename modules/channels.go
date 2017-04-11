/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    CHANNEL_MODULE = 0x20000000
    CHANNEL_CONFIG = MAIN_CONFIG|CONFIG_BLOCK
)

var channel = String{ len("channel"), "channel" }
var channelContext = &Context{
    channel,
    nil,
    nil,
}

var channels = String{ len("channels"), "channels" }
var channelCommands = []Command{

    { channels,
      CHANNEL_CONFIG,
      channelsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func channelsBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    if nil == cycle {
        cycle.Error("channels block error")
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    cycle.Block(cycle, CHANNEL_MODULE, flag)

    return Ok
}

var channelModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(channelContext),
    channelCommands,
    CONFIG_MODULE,
    nil,
    nil,
}

func init() {
    Modules = Load(Modules, &channelModule)
}