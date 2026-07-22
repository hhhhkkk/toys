package main

import (
	"fmt"
	"slices"
)

type HuaSe int

const (
	FangPian = HuaSe(1)
	MeiHua   = HuaSe(2)
	HongTao  = HuaSe(3)
	HeiTao   = HuaSe(4)
	Xiaowang = HuaSe(5)
	Dawang   = HuaSe(6)
)

func (h HuaSe) ToString() (ret string) {
	switch {
	case h == HeiTao:
		ret = "黑桃"
	case h == HongTao:
		ret = "红桃"
	case h == MeiHua:
		ret = "梅花"
	case h == FangPian:
		ret = "方片"
	case h == Xiaowang:
		ret = "小王"
	case h == Dawang:
		ret = "大王"
	}
	return ret
}

func (h HuaSe) Weight() (ret int) {
	switch {
	case h == HeiTao:
		ret = 4
	case h == HongTao:
		ret = 3
	case h == MeiHua:
		ret = 2
	case h == FangPian:
		ret = 1
	case h == Xiaowang:
		ret = 10000
	case h == Dawang:
		ret = 20000
	}
	return ret
}

type Operator int

func (o Operator) ToString() (ret string) {
	switch o {
	case 1:
		ret = "A"
	case 2:
		ret = "2"
	case 3:
		ret = "3"
	case 4:
		ret = "4"
	case 5:
		ret = "5"
	case 6:
		ret = "6"
	case 7:
		ret = "7"
	case 8:
		ret = "8"
	case 9:
		ret = "9"
	case 10:
		ret = "10"
	case 11:
		ret = "J"
	case 12:
		ret = "Q"
	case 13:
		ret = "K"
	}
	return ret
}

func (o Operator) Weight() (ret int) {
	switch o {
	case 1:
		ret = 2000
	case 2:
		ret = 3000
	case 3:
		ret = 300
	case 4:
		ret = 400
	case 5:
		ret = 500
	case 6:
		ret = 600
	case 7:
		ret = 700
	case 8:
		ret = 800
	case 9:
		ret = 900
	case 10:
		ret = 1000
	case 11:
		ret = 1100
	case 12:
		ret = 1200
	case 13:
		ret = 1300
	}
	return ret
}

type pai struct {
	Number Operator
	HuaSe  HuaSe
}

func NewPai(h HuaSe, n int) pai {
	return pai{
		HuaSe:  h,
		Number: Operator(n),
	}
}

func (p pai) Print() {
	fmt.Printf("%d %d\n", p.HuaSe, p.Number)
}

func (p pai) ToString() string {
	if slices.Contains([]string{"大王", "小王"}, p.HuaSe.ToString()) {
		return p.HuaSe.ToString()
	}
	return fmt.Sprintf("%s %s", p.HuaSe.ToString(), p.Number.ToString())
}

func (p pai) Weight() int {
	return p.HuaSe.Weight() + p.Number.Weight()
}
