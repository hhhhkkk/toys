package main

import (
	"fmt"
)

const (
	HeiTao   = HuaSe(1)
	HongTao  = HuaSe(2)
	MeiHua   = HuaSe(3)
	FangPian = HuaSe(4)
	Xiaowang = HuaSe(5)
	Dawang   = HuaSe(6)
)

type HuaSe int

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

type pai struct {
	Number Operator
	HuaSe  HuaSe
}

func (p *pai) Print() {
	fmt.Printf("%d %d\n", p.HuaSe, p.Number)
}

func (p *pai) ToString() string {
	return fmt.Sprintf("%s %s", p.HuaSe.ToString(), p.Number.ToString())
}

func main() {
	// 准备牌集合
	game := NewDouDiZhu()

	isDouDiZhi, dizhu := game.(*DouDiZhu)
	if !dizhu {
		panic("err")
	}
	// isDouDiZhi.PrintAll()
	isDouDiZhi.Shuffle()
	isDouDiZhi.Begin()
}
