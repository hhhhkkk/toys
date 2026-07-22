package main

import "testing"

func TestCollection(t *testing.T) {
	// 测试是否按序执行
	pais := []pai{
		{
			HuaSe: Xiaowang,
		},
		{
			HuaSe: Dawang,
		},
		{
			HuaSe:  HeiTao,
			Number: Operator(1),
		},
		{
			HuaSe:  HongTao,
			Number: Operator(2),
		},
		{
			HuaSe:  FangPian,
			Number: Operator(2),
		},
	}

	c := NewCollection()
	c.pais = pais
	c.Sort()

	resultStr := "大王小王红桃 2方片 2黑桃 A"
	ret := ""
	for _, p := range c.pais {
		ret += p.ToString()
	}

	if resultStr != ret {
		t.Errorf("错误[%s]", ret)
	}
}
