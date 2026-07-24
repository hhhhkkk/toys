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

func (t ThreeWith) Check(c *Collection) bool {
	l := len(c.pais)
	// 只可能是 3、4、5 张牌
	if l > 5 || l < 2 {
		return false
	}
	c.Sort()
	// 如果是三不带，则三张应该一样
	if l == 3 {
		return isSame(c)
	}
	// 如果是三带一， 那么第二张或者第三张一定是三个的牌
	// 如果是三带二，那么第 三中一定是三个的牌
	// 取第三张作为基准牌
	base := c.pais[2]
	diff := slices.IndexFunc(c.pais, func(p pai) bool {
		return p.Number != base.Number
	})

	// 牌完全一样
	if diff == -1 {
		return false
	}

	// 三带排完序后，一定是以下情况
	// 要么前一、二张与后三张不一样
	// 要么后一、二张与后三张不一样
	// 找到是前缀还是后缀， 然后对比缀是不是一样的
	// 前缀
	suffix := NewCollection()
	suffix.pais = c.pais[:3]
	if isSame(suffix) {
		prefix := NewCollection()
		prefix.pais = c.pais[3:]
		return isSame(prefix)
	}

	prefix := NewCollection()
	prefix.pais = c.pais[l-3:]
	if isSame(prefix) {
		nsuffix := NewCollection()
		if l == 4 {
			nsuffix.pais = c.pais[:1]
		}
		if l == 5 {
			nsuffix.pais = c.pais[:2]
		}
		return isSame(nsuffix)
	}
	return false
}

// Plane 飞机.
type Plane struct {
	subRule Three
}

var _ Rule = (*Plane)(nil)

func (p Plane) Check(c *Collection) (ret bool) {
	ret = false
	// 顺子通常是 n 个 三带一、 n 个三带二、n 个三不带
	l := len(c.pais)
	if l > 20 || l < 6 {
		return
	}

	c.Sort()
	// 3 不带：6、9、12、15、18
	// 3 带 1：8、12、16、20
	// 3 带 2：10、15、20

	// 统计出来 3 个张数 机体，如果有一个是大于 3 的，就不是
	// 判断所有 3 的个数是否是连续的
	// 判断翅膀， 其应该是 机体个数 相等或者是 2n,否则报错
	// 如果是 2n , 则判断是否都是对， 否则报错
	m := make(map[int]int)

	for _, v := range c.pais {
		if _, ok := m[int(v.Number)]; !ok {
			m[int(v.Number)] = 1
		} else {
			m[int(v.Number)] += 1
		}
	}

	mainPais := make([]pai, 0)
	attchPais := make([]pai, 0)
	// 有超过 3 的
	for k, v := range m {
		if v > 3 {
			return
		}
		mainPai := slices.DeleteFunc(c.pais, func(p pai) bool {
			return int(p.Number) == k
		})
		if v == 3 {
			mainPais = append(mainPais, mainPai...)
			continue
		}
		attchPais = append(attchPais, mainPai...)
	}

	// 所有主牌都是连续的
	basePai := mainPais[0]
	for _, v := range mainPais[1:] {
		if v.Number != basePai.Number && v.Number != basePai.Number+1 {
			return
		}
	}

	ml := len(m)
	al := len(attchPais)
	if al%ml != 0 {
		return false
	}

	// 都是对
	// 都是单无需验证
	if al > ml {
		tmpC := NewCollection()
		tmpC.pais = attchPais
		tmpC.Sort()

		chunk := slices.Chunk(tmpC.pais, 2)
		dui := Dui{}
		for v := range chunk {
			tc := NewCollection()
			tc.pais = v
			if dui.Check(tc) != true {
				return false
			}
		}
	}
	ret = true
	return ret
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
