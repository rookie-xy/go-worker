
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type Option struct {
    File    string

    Action  interface {
        Set()
        Get()
    }
}
