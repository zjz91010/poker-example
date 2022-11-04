package main

import (
	"doudizhutest/computer"
	"fmt"
)

func main() {
	//测试构造一副牌，洗牌，发牌
	createNewShuffle()

	//根据牌面数量判断牌面类型
	playing()
}

//根据牌面数量判断牌面类型
func playing() {
	// cardsA := []string{"A3", "K17"}
	cardsA := []string{"A3", "B3", "C3", "A4", "B4", "C4", "A5", "B5", "A5", "A6", "B6", "A6", "A11", "A7", "B12", "B7"}
	playA := computer.ParseCardsInSize(cardsA)
	fmt.Printf("playA：%v\n", playA)

	cardsB := []string{"A4", "B4", "C4", "A5", "B5", "C5", "A6", "B6", "A6", "A7", "B7", "A7", "A11", "A10", "B12", "13"}
	playB := computer.ParseCardsInSize(cardsB)
	fmt.Printf("playB：%v\n", playB)
}

//测试构造一副牌，洗牌，发牌
func createNewShuffle() {
	// fmt.Println("main测试")
	initValues := computer.CreateNew()
	fmt.Printf("构造一副牌：%v\n", initValues)

	computer.Shuffle(initValues)
	// fmt.Printf("洗牌后：%v\n", initValues)

	nikename := "玩家"
	var playCards []string
	for i := 0; i <= 3; i++ {
		if i == 3 {
			nikename = "底牌"
		}
		playCards = computer.Dispacther(i, initValues)
		fmt.Printf("%s%d：%v\n", nikename, i, playCards)
	}

}
