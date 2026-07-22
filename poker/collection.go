package main

import (
	"fmt"
	"slices"
)

type SortHandle func(p1, p2 pai) int

type Collection struct {
	pais       []pai
	SortHandle SortHandle
}

type Option func(c *Collection)

func (c *Collection) Print() {
	for _, p := range c.pais {
		fmt.Println(p.ToString())
	}
}

func NewCollection(opt ...Option) *Collection {
	collec := &Collection{
		pais:       make([]pai, 0),
		SortHandle: DefaultSortHandle,
	}
	for _, f := range opt {
		f(collec)
	}
	return collec
}

func DefaultSortHandle(p1, p2 pai) int {
	// 先比大小王
	if p1.HuaSe == Dawang {
		return -1
	}
	if p2.HuaSe == Dawang {
		return 1
	}

	if p1.HuaSe == Xiaowang {
		return -1
	}

	if p2.HuaSe == Xiaowang {
		return 1
	}
	// 比数字
	if p1.Weight() > p2.Weight() {
		return -1
	}
	if p1.Weight() < p2.Weight() {
		return 1
	}
	// 比花色
	if p1.HuaSe > p2.HuaSe {
		return -1
	} else if p1.HuaSe < p2.HuaSe {
		return 1
	} else {
		return 0
	}
}

func (c *Collection) Sort() *Collection {
	slices.SortFunc(c.pais, c.SortHandle)
	return c
}
