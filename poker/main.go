package main

func main() {
	var players []*Player
	// 准备牌集合
	game := NewDouDiZhu(players)

	isDouDiZhi, dizhu := game.(*DouDiZhu)
	if !dizhu {
		panic("err")
	}
	// isDouDiZhi.PrintAll()
	isDouDiZhi.Shuffle()
	isDouDiZhi.Begin()
}
