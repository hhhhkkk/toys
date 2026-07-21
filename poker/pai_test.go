package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestPaiHuase(t *testing.T) {
	tests := []struct {
		name string
		p    pai
	}{
		{
			name: "黑桃",
			p: pai{
				Number: Operator(1),
				HuaSe:  HeiTao,
			},
		},
		{
			name: "红桃",
			p: pai{
				Number: Operator(1),
				HuaSe:  HongTao,
			},
		},
		{
			name: "梅花",
			p: pai{
				Number: Operator(1),
				HuaSe:  MeiHua,
			},
		},
		{
			name: "方片",
			p: pai{
				Number: Operator(1),
				HuaSe:  FangPian,
			},
		},
		{
			name: "小王",
			p: pai{
				Number: Operator(1),
				HuaSe:  Xiaowang,
			},
		},
		{
			name: "大王",
			p: pai{
				Number: Operator(1),
				HuaSe:  Dawang,
			},
		},
	}

	for _, tu := range tests {
		t.Run(tu.name, func(t *testing.T) {
			if tu.name != tu.p.HuaSe.ToString() {
				t.Error("牌对不上号")
			}
		})
	}
	// 花色大小
	ps := make([]pai, 0, len(tests))

	for _, tu := range tests {
		ps = append(ps, tu.p)
	}
	slices.SortFunc(ps, func(p1, p2 pai) int {
		if p1.HuaSe > p2.HuaSe {
			return -1
		} else if p1.HuaSe < p2.HuaSe {
			return 1
		}
		return 0
	})

	resultString := "大王小王黑桃红桃梅花方片"
	var result string
	for _, p := range ps {
		result += p.HuaSe.ToString()
	}
	t.Run("花色比大小", func(t *testing.T) {
		fmt.Println(result)
		if resultString != result {
			t.Error("花色排序不对")
		}
	})
}

func TestPaiNumber(t *testing.T) {
	tests := []pai{
		{
			Number: Operator(1),
			HuaSe:  HeiTao,
		},
		{
			Number: Operator(2),
			HuaSe:  HeiTao,
		},
		{
			Number: Operator(3),
			HuaSe:  HeiTao,
		},
		{
			Number: Operator(4),
			HuaSe:  HeiTao,
		},
		{
			Number: Operator(1),
			HuaSe:  Xiaowang,
		},
		{
			Number: Operator(1),
			HuaSe:  Dawang,
		},
	}

	slices.SortFunc(tests, func(p1, p2 pai) int {
		if p1.HuaSe == Dawang {
			return -1
		}
		if p1.HuaSe == Xiaowang {
			if p2.HuaSe == Dawang {
				return 1
			} else {
				return -1
			}
		}
		if p1.Number > p2.Number {
			return -1
		} else if p1.Number < p2.Number {
			return 1
		}
		return 0
	})

	var str string
	for _, p := range tests {
		str += p.ToString()
	}
	var rightString = "大王小王黑桃 4黑桃 3黑桃 2黑桃 A"
	if rightString != str {
		t.Errorf("对不上: %s", str)
	}
}
