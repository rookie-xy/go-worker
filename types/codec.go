/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type Codec struct {
    name   string
    codec  CodecIf
}

type CodecIf interface {
    Encode()
    Decode()
}

func NewCodec() *Codec {
    return &Codec{}
}
