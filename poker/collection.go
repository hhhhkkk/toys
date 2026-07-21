package main

import (
	"fmt"
	"slices"
)

type Collection struct {
	pais []pai
}

type Option func(c *Collection)

func (c *Collection) Print() {
	for _, p := range c.pais {
		fmt.Println(p.ToString())
	}
}

func NewCollection(opt ...Option) *Collection {
	collec := &Collection{
		pais: make([]pai, 0),
	}
	for _, f := range opt {
		f(collec)
	}
	return collec
}

func (c *Collection) Sort(isAsc bool) *Collection {
	var gt, lt int
	if isAsc {
		gt, lt = 1, -1
	} else {
		gt, lt = -1, 1
	}
	slices.SortFunc(c.pais, func(p1, p2 pai) int {
		if p1.HuaSe > p2.HuaSe {
			return gt
		} else if p1.HuaSe < p2.HuaSe {
			return lt
		} else {
			return 0
		}
	})
	return c
}
