/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type Channal struct {
    *Cycle
    *File

     upstream   chan string
     downstream chan string

     channal    ChannalIf
}

type ChannalIf interface {
    push(name string, data string) int
    pull(name string) int
}

func NewChannal() *Channal {
    return &Channal{}
}

func (c *Channal) push(name string, data string) int {
    c.upstream <- data

    return Ok
}

func (c *Channal) pull(name string) string {
    data := <-c.downstream

    return data
}
