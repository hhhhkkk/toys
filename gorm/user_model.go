package main

import (
	"context"
	"fmt"
	"math/rand/v2"

	"gorm.io/gorm"
)

type User struct {
	Id   int
	Name string
	db   *gorm.DB
}

func Run(db *gorm.DB) {
	u := NewUser()

	u.Name = ("admin" + randString(6))

	ctx := context.Background()

	// 自动建表
	db.AutoMigrate(&User{})

	whatsThis := gorm.G[User](db)

	err := whatsThis.Create(ctx, u)
	if err != nil {
		panic("创建用户失败")
	}
	fmt.Printf("create user[%d], %s\n", u.Id, u.Name)

	user, err := whatsThis.Where("id = ?", 5).First(ctx)
	if err != nil {
		panic("user [1] not found")
	}
	fmt.Printf("id %d, name %s\n", user.Id, user.Name)

	count, err := whatsThis.Count(ctx, "*")
	if err != nil {
		fmt.Printf("总数 [%d]\n", count)
	}

	randUpdateId := rand.IntN(int(count))
	eff, err := whatsThis.Where("id = ?", randUpdateId).Update(ctx, "name", gorm.Expr("concat(name, 'manual')"))

	if err != nil {
		panic("update fail")
	}

	if eff == 0 {
		panic(fmt.Errorf("update not found [%d]", randUpdateId))
	}

	// delete

	delId := 100000
	eff, err = whatsThis.Where("id = ?", delId).Delete(ctx)
	if err != nil {
		panic("delete fail")
	}

	if eff == 0 {
		fmt.Errorf("delete not found [%d]", randUpdateId)
	}
}

func NewUser() *User {
	return &User{}
}

func randString(n int) string {
	bases := []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'k',
		'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '~', '!', '#', '$', '%', '0', '1',
		'2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C',
		'D', 'E', 'F', 'G',
	}
	basesLen := len(bases)
	m := make(map[int]bool)
	ret := []byte{}
	for {
		if len(ret) == n {
			break
		}
		randIndex := rand.IntN(basesLen)
		if _, ok := m[randIndex]; ok {
			continue
		}
		ret = append(ret, bases[randIndex])
		m[randIndex] = true
	}
	return string(ret)
}
