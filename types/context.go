
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type CreateConfFunc func(cycle *Cycle) Void;
type InitConfFunc func(cycle *Cycle, conf *Void) string;


type Context struct {
    Name  String

    Conf  interface {
        Create(cycle *Cycle)
        Init(cycle *Cycle)
    }

    //CreateConf   CreateConfFunc
    //InitConf     InitConfFunc
}
