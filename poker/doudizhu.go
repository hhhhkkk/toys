package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"sync"
)

type Player struct {
	Name string
	col  *Collection
}

type Game interface {
	// GetAllPai() *Collection
	Shuffle()
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

	// player
	Playeries []*Player

	// 手牌

	// passed *
	mu sync.RWMutex
}

func NewDouDiZhu(p []*Player) Game {
	d := &DouDiZhu{
		col:       NewCollection(),
		Playeries: p,
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
			var value int
			// A and 2
			if i == 0 || i == 1 {
				value = 10 + i
			} else {
				value = i + 1
			}

			r.col.pais = append(r.col.pais, pai{
				Number: Operator(value),
				HuaSe:  j,
			})
		}
	}
	// 大小王
	r.col.pais = append(r.col.pais, pai{
		Number: Operator(20),
		HuaSe:  Xiaowang,
	}, pai{
		Number: Operator(21),
		HuaSe:  Dawang,
	})
}

func (r *DouDiZhu) PrintAll() {
	r.col.Print()
}

func (r *DouDiZhu) Shuffle() {
	r.mu.Lock()
	defer r.mu.Unlock()

	rand.Shuffle(len(r.col.pais), func(i, j int) {
		r.col.pais[i], r.col.pais[j] = r.col.pais[j], r.col.pais[i]
	})
}

var _ Game = (*DouDiZhu)(nil)

func (r *DouDiZhu) Begin() {
	playeries := []Player{}

	r.mu.Lock()
	defer r.mu.Unlock()
	for range playeries {
		fmt.Println(len(r.col.pais))

		fenpai := slices.Clone(r.col.pais[:17])
		r.col.pais = r.col.pais[17:]
		c := &Collection{
			pais: fenpai,
		}
		c.Print()
	}

	lastC := &Collection{
		pais: r.col.pais,
	}
	lastC.Print()
}
