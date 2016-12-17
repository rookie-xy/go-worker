
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


var (
    WkrOk    =  0
    WkrError = -1
    WkrAgain = -2
)


var (
    LF   =  byte('\x0a')
    CR   =  byte('\x0d')
    CRLF =  "\x0d\x0a"
)


type InitMoudleFunc func(cycle *DvrCycleType) uint;
type InitProcessFunc func(cycle *DvrCycleType) uint;
type CreateConfFunc func(cycle *DvrCycleType) DvrVoidType;
type InitConfFunc func(cycle *DvrCycleType, conf *DvrVoidType) string;
type SetFunc func(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string;


type Context struct {
}


type Command struct {
    Name    StrType;
    Type    uint;
    Set     SetFunc;
    Conf    int;
    Offset  uintptr;
    Post    InterfType;
};


type Module struct {
    CtxIndex      uint;
    Index         uint;
    Ctx          *CoreModule;
    Commands      []Command;
    Type          uint;
    InitMoudle   *InitMoudleFunc;
    InitProcess  *InitProcessFunc;
}
