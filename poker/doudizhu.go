package main

import (
	"math/rand/v2"
	"slices"
	"sync"
)

type Player struct {
	Name    string
	isDiZhu bool
	Col     *Collection
}

type Game interface {
	// GetAllPai() *Collection
	Begin()
	// Pass() // 出牌
	// PassRecord() // 有序 slice
}

type DouDiZhu struct {
	// 所有牌
	col *Collection

	// 出过的牌

	// 剩余的牌

	// 暗牌
	dizhuPai *Collection

	// player
	Playeries []*Player

	// passed *
	mu sync.RWMutex
}

func NewDouDiZhu(p []*Player) Game {
	d := &DouDiZhu{
		col:       NewCollection(),
		Playeries: p,
		dizhuPai:  NewCollection(),
	}
	d.initPai()
	return d
}

func (r *DouDiZhu) initPai() {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 字母 花色
	for i := range 13 {
		for _, j := range []HuaSe{HeiTao, HongTao, MeiHua, FangPian} {

			r.col.pais = append(r.col.pais, pai{
				Number: Operator(i + 1),
				HuaSe:  j,
			})
		}
	}
	// 大小王
	r.col.pais = append(r.col.pais, pai{
		Number: Operator(30),
		HuaSe:  Xiaowang,
	}, pai{
		Number: Operator(30),
		HuaSe:  Dawang,
	})

	// 打乱
	rand.Shuffle(len(r.col.pais), func(i, j int) {
		r.col.pais[i], r.col.pais[j] = r.col.pais[j], r.col.pais[i]
	})
}

func (r *DouDiZhu) PrintAll() {
	r.col.Print()
}

var _ Game = (*DouDiZhu)(nil)

func (r *DouDiZhu) Begin() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, p := range r.Playeries {
		fenpai := slices.Clone(r.col.pais[:17])
		r.col.pais = r.col.pais[17:]
		c := NewCollection()
		c.pais = fenpai
		p.Col = c.Sort()
	}
	r.dizhuPai = NewCollection()
	r.dizhuPai.pais = r.col.pais
}

func (r *DouDiZhu) CallDiZhu() {
	r.mu.Lock()
	defer r.mu.Unlock()
	dizhu := rand.IntN(len(r.Playeries))

	p := r.Playeries[dizhu]
	if p == nil {
		panic("玩家不存在")
	}
	p.Col.pais = append(p.Col.pais, r.dizhuPai.pais...)
	p.Col.Sort()
}
