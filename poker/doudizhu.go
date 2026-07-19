package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"sync"
)

type Player int

type Game interface {
	// GetAllPai() *Collection
	Shuffle()
	Begin()
	// Pass() // 出牌
	// PassRecord() // 有序 slice
}

type DouDiZhu struct {
	col *Collection
	// passed *
	mu sync.RWMutex
}

func NewDouDiZhu() Game {
	d := &DouDiZhu{
		col: NewCollection(),
	}
	d.init()
	return d
}

func (r *DouDiZhu) init() {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 字母 花色
	for i := range 13 {
		for _, j := range []HuaSe{HeiTao, HongTao, MeiHua, FangPian} {
			r.col.pais = append(r.col.pais, &pai{
				Number: Operator(i + 1),
				HuaSe:  j,
			})
		}
	}
	// 大小王
	r.col.pais = append(r.col.pais, &pai{
		Number: Operator(14),
		HuaSe:  Xiaowang,
	}, &pai{
		Number: Operator(15),
		HuaSe:  Dawang,
	})
}

func (r *DouDiZhu) GetAllPai() *Collection {
	return r.col
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

type Collection struct {
	pais []*pai
}

type Option func(c *Collection)

func (c *Collection) Print() {
	for _, p := range c.pais {
		fmt.Println(p.ToString())
	}
}

func NewCollection(opt ...Option) *Collection {
	collec := &Collection{
		pais: make([]*pai, 0),
	}
	for _, f := range opt {
		f(collec)
	}
	return collec
}

func (r *DouDiZhu) Begin() {
	playeries := []Player{Player(1), Player(2), Player(3)}

	r.mu.Lock()
	defer r.mu.Unlock()
	for _, player := range playeries {
		fmt.Println(len(r.col.pais))

		fenpai := slices.Clone(r.col.pais[:17])
		r.col.pais = r.col.pais[17:]
		c := &Collection{
			pais: fenpai,
		}
		fmt.Printf("%d, %d\n", player, len(fenpai))
		c.Print()
	}

	lastC := &Collection{
		pais: r.col.pais,
	}
	fmt.Println("last")
	lastC.Print()
}
