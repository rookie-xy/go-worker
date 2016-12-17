
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type Routine struct {
    File    string

    Action  interface {
        Set()
        Get()
        Parse()
    }
}
