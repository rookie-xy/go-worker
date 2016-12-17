
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type Signal struct {
    File    string

    Action  interface {
        Set()
        Get()
        Parse()
    }
}
