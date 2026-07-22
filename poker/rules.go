package main

import (
	"fmt"
	"slices"
)

// Rule 牌型
type Rule interface {
	Check(c *Collection) bool
}

var _ Rule = (*Boom)(nil)

type Boom struct{}

func (b Boom) Check(c *Collection) bool {
	l := len(c.pais)
	if l != 2 && l != 4 {
		return false
	}

	if l == 2 {
		for _, p := range c.pais {
			if p.HuaSe != Dawang && p.HuaSe != Xiaowang {
				return false
			}
		}
		return true
	}
	return isSame(c)
}

type Single struct{}

var _ Rule = (*Single)(nil)

func (s Single) Check(c *Collection) bool {
	return len(c.pais) == 1
}

type Dui struct{}

var _ Rule = (*Dui)(nil)

func (s Dui) Check(c *Collection) bool {
	if len(c.pais) != 2 {
		return false
	}
	return isSame(c)
}

type Three struct{}

var _ Rule = (*Three)(nil)

func (s Three) Check(c *Collection) bool {
	if len(c.pais) != 3 {
		return false
	}
	return isSame(c)
}

// ThreeWith 三带.
type ThreeWith struct {
	subRule Three
}

// Plane 飞机.
type Plane struct {
	subRule Three
}

// Straight 顺子.
type Straight struct{}

// StraightDui 连对.
type StraightDui struct {
	subRule  Dui
	subRule2 Straight
}

// FourWith 四带二
type FourWith struct{}

func isSame(c *Collection) bool {
	l := len(c.pais)
	if l == 0 {
		panic("no pai")
	}
	if l < 2 {
		return true
	}

	first := c.pais[0]
	var zeroNum Operator
	for _, v := range c.pais[1:] {
		if v.Number == zeroNum || v.Number != first.Number {
			return false
		}
	}
	return true
}

func isIndex(c *Collection) bool {
	if len(c.pais) < 3 {
		return false
	}
	forbiddenNumber := Operator(2)
	forbiddenHuase := []HuaSe{Dawang, Xiaowang}
	for _, v := range c.pais {
		if v.Number == forbiddenNumber || slices.Contains(forbiddenHuase, v.HuaSe) {
			return false
		}
	}
	c.Sort()
	first := c.pais[0].Number
	for k, v := range c.pais[1:] {
		if first == Operator(1) {
			first = Operator(13)
		} else {
			first -= 1
		}
		fmt.Println(k, v.Number, first)
		if v.Number != first {
			return false
		}
	}
	return true
}
