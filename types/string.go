
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type String struct {
    Len int;
    Data String;
}


NilString = String{ 0, nil };
