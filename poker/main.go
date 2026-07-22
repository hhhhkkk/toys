package main

func main() {
	players := []*Player{
		{
			Name: "张三",
			Col:  NewCollection(),
		},
		{
			Name: "李四",
		},
		{
			Name: "王二",
		},
	}
	// 准备牌集合
	game := NewDouDiZhu(players)

	isDouDiZhi, dizhu := game.(*DouDiZhu)
	if !dizhu {
		panic("err")
	}

	isDouDiZhi.Begin()
	isDouDiZhi.CallDiZhu()
}
