package main

import (
	"testing"
)

func TestBoom(t *testing.T) {
	units := []struct {
		p      []pai
		name   string
		rule   Rule
		result bool
	}{
		{
			name:   "王炸",
			rule:   Boom{},
			result: true,
			p: []pai{
				{
					HuaSe: Dawang,
				},
				{
					HuaSe: Xiaowang,
				},
			},
		},
		{
			name:   "非王炸",
			rule:   Boom{},
			result: false,
			p: []pai{
				{
					HuaSe:  HeiTao,
					Number: Operator(1),
				},
				{
					HuaSe: Xiaowang,
				},
			},
		},
		{
			name:   "4个A",
			rule:   Boom{},
			result: true,
			p: []pai{
				{
					HuaSe:  HeiTao,
					Number: Operator(1),
				},
				{
					HuaSe:  HongTao,
					Number: Operator(1),
				}, {
					HuaSe:  MeiHua,
					Number: Operator(1),
				}, {
					HuaSe:  FangPian,
					Number: Operator(1),
				},
			},
		},
		{
			name:   "3个21个A",
			rule:   Boom{},
			result: false,
			p: []pai{
				{
					HuaSe:  HeiTao,
					Number: Operator(2),
				},
				{
					HuaSe:  HongTao,
					Number: Operator(2),
				}, {
					HuaSe:  MeiHua,
					Number: Operator(2),
				}, {
					HuaSe:  FangPian,
					Number: Operator(1),
				},
			},
		},
	}

	for _, uint := range units {
		col := NewCollection()
		col.pais = uint.p
		res := uint.rule.Check(col)
		if uint.result != res {
			t.Error(uint.name + "fail")
		} else {
			s := ""
			for _, v := range uint.p {
				s += v.ToString()
			}
			t.Logf("%s [%s] %t", uint.name, s, res)
		}
	}
}

func TestIsSame(t *testing.T) {
	units := []struct {
		name   string
		p      []pai
		rule   Rule
		result bool
	}{
		{
			name: "一张牌 -> 两张牌",
			p: []pai{
				{
					HuaSe: Dawang,
				},
				{
					HuaSe: Xiaowang,
				},
			},
			rule:   Single{},
			result: false,
		},
		{
			name: "一张牌 -> 一张牌",
			p: []pai{
				{
					HuaSe: Dawang,
				},
			},
			rule:   Single{},
			result: true,
		},
		{
			name: "两张牌 -> 大小王",
			p: []pai{
				{
					HuaSe:  Dawang,
					Number: Operator(30),
				},
				{
					HuaSe:  Dawang,
					Number: Operator(20),
				},
			},
			rule:   Dui{},
			result: false,
		},
		{
			name: "两张牌 -> 对A",
			p: []pai{
				{
					HuaSe:  HeiTao,
					Number: Operator(1),
				},
				{
					HuaSe:  HongTao,
					Number: Operator(1),
				},
			}, rule: Dui{},
			result: true,
		},
		{
			name: "两张牌 -> 1A12",
			p: []pai{
				{
					HuaSe:  HeiTao,
					Number: Operator(1),
				},
				{
					HuaSe:  HongTao,
					Number: Operator(2),
				},
			},
			rule:   Dui{},
			result: false,
		},
	}

	for _, uint := range units {
		col := NewCollection()
		col.pais = uint.p
		res := uint.rule.Check(col)
		if uint.result != res {
			t.Fatal(uint.name + "fail")
		} else {
			s := ""
			for _, v := range uint.p {
				s += v.ToString()
			}
			t.Logf("%s [%s] %t", uint.name, s, res)
		}
	}
}

func TestIsIndex(t *testing.T) {
	// 包含 2
	p1 := []pai{
		{HuaSe: HeiTao, Number: Operator(2)},
		{HuaSe: HeiTao, Number: Operator(3)},
		{HuaSe: HeiTao, Number: Operator(4)},
		{HuaSe: HeiTao, Number: Operator(5)},
		{HuaSe: HeiTao, Number: Operator(6)},
	}

	p2 := []pai{
		{HuaSe: Dawang},
		{HuaSe: HeiTao, Number: Operator(3)},
		{HuaSe: HeiTao, Number: Operator(4)},
		{HuaSe: HeiTao, Number: Operator(5)},
		{HuaSe: HeiTao, Number: Operator(6)},
	}

	p3 := []pai{
		{HuaSe: Xiaowang},
		{HuaSe: HeiTao, Number: Operator(3)},
		{HuaSe: HeiTao, Number: Operator(4)},
		{HuaSe: HeiTao, Number: Operator(5)},
		{HuaSe: HeiTao, Number: Operator(6)},
	}

	p4 := []pai{
		{HuaSe: HeiTao, Number: Operator(3)},
		{HuaSe: HeiTao, Number: Operator(4)},
		{HuaSe: HeiTao, Number: Operator(5)},
		{HuaSe: HeiTao, Number: Operator(6)},
		{HuaSe: HeiTao, Number: Operator(8)},
	}

	p5 := []pai{
		{HuaSe: HeiTao, Number: Operator(3)},
		{HuaSe: HeiTao, Number: Operator(3)},
		{HuaSe: HongTao, Number: Operator(5)},
		{HuaSe: HeiTao, Number: Operator(6)},
		{HuaSe: HeiTao, Number: Operator(7)},
	}

	p6 := []pai{
		{HuaSe: HeiTao, Number: Operator(3)},
		{HuaSe: HeiTao, Number: Operator(4)},
		{HuaSe: HongTao, Number: Operator(5)},
		{HuaSe: HeiTao, Number: Operator(6)},
		{HuaSe: HeiTao, Number: Operator(7)},
		{HuaSe: FangPian, Number: Operator(8)},
		{HuaSe: HeiTao, Number: Operator(9)},
		{HuaSe: HeiTao, Number: Operator(10)},
		{HuaSe: MeiHua, Number: Operator(11)},
		{HuaSe: HeiTao, Number: Operator(12)},
		{HuaSe: HeiTao, Number: Operator(13)},
		{HuaSe: HeiTao, Number: Operator(13)},
		{HuaSe: HeiTao, Number: Operator(1)},
	}

	p7 := []pai{
		{HuaSe: HeiTao, Number: Operator(3)},
		{HuaSe: HeiTao, Number: Operator(4)},
		{HuaSe: HongTao, Number: Operator(5)},
		{HuaSe: HeiTao, Number: Operator(6)},
		{HuaSe: HeiTao, Number: Operator(7)},
	}
	n := [][]pai{
		p1, p2, p3, p4, p5, p6, p7,
		// p7,
	}
	for k, v := range n {
		c := NewCollection()
		c.pais = v
		if k <= 5 && isIndex(c) != false {
			t.Fatalf("P%d", k+1)
		}
		if k > 5 && isIndex(c) != true {
			t.Fatalf("P%d", k+1)
		}
	}
}

func TestThreeWith(t *testing.T) {
	tests := [][]pai{
		{
			NewPai(HeiTao, 1),
			NewPai(HeiTao, 1),
			NewPai(HeiTao, 1),
		},
		{
			NewPai(HeiTao, 3),
			NewPai(HeiTao, 4),
			NewPai(HeiTao, 5),
		},
		{
			NewPai(Dawang, 0),
			NewPai(HeiTao, 1),
			NewPai(HeiTao, 1),
			NewPai(HeiTao, 1),
		},
		{
			NewPai(Xiaowang, 0),
			NewPai(HeiTao, 1),
			NewPai(HeiTao, 1),
			NewPai(HeiTao, 1),
		},
		{
			NewPai(HongTao, 2),
			NewPai(HeiTao, 2),
			NewPai(HeiTao, 1),
			NewPai(HeiTao, 1),
		},
		{
			NewPai(HongTao, 4),
			NewPai(HeiTao, 5),
			NewPai(HeiTao, 6),
			NewPai(HeiTao, 7),
			NewPai(HeiTao, 8),
		},
		{
			NewPai(HongTao, 2),
			NewPai(HeiTao, 2),
			NewPai(HeiTao, 3),
			NewPai(HeiTao, 3),
			NewPai(HeiTao, 3),
		},
	}

	handler := ThreeWith{}
	for k, p := range tests {
		c := NewCollection()
		c.pais = p
		res := []bool{true, false, true, true, false, false, true}
		if handler.Check(c) != res[k] {
			t.Errorf("p%d fail", k)
			return
		} else {
			t.Logf("p%d pass", k)
		}
	}
}
