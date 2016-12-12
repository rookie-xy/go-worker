
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type Configure struct {
    File    string

    Action  interface {
        Set()
        Get()
        Parse()
    }
}
