
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type Context struct {
    Name  String

    Data  interface {
        Create(cycle *Cycle)
        Init(cycle *Cycle)
    }
}
