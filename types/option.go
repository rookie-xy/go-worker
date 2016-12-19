
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type Option struct {
    File    string

    Data  interface {
        Set(argc int, argv []string) int
        Get(argc int, argv []string) int
    }
}
