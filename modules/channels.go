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
    CHANNEL_CONFIG = 0x02000000
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
      MAIN_CONFIG|CONFIG_MAP,
      channelsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func channelsBlock(cycle *Cycle, _ *Command, _ *unsafe.Pointer) int {
    cycle.Configure.Block(CHANNEL_MODULE, CHANNEL_CONFIG)
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
    //Modules = Load(Modules, &channelModule)
}