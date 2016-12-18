
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type SetFunc func(cf *Configure, cmd *Command, conf *Void) string;


type Command struct {
    Name    String;
    Type    uint;
    Set     SetFunc;
    Conf    int;
    Offset  uintptr;
    Post    InterfType;
};


NilCommand = Command{ NilString, 0, nil, 0, 0, nil };
