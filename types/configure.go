
/*
 * Copyright (C) 2016 Meng Shi
 */


package types 


type Configure struct {
    I      int
    Action interface {
        Set()
        Get()
        Parse()
    }
}
