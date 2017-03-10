/*
 * Copyright (C) 2017 Meng Shi
 */

package types

type Filter struct {
    name    string
    filter  FilterIf
}

type FilterIf interface {
    Filter()
}

func NewFilter() *Filter {
    return &Filter{}
}
