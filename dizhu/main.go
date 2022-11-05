package main

import (
	"dizhu/computer"
	"dizhu/util"
	"fmt"
)

func main() {
	// fmt.Printf("%v\n", "好牌")

	newCards := computer.CreateNewCards()
	fmt.Printf("构造一副牌:%v, %d张\n", newCards, len(newCards))
	computer.Shuffle(newCards)
	// fmt.Printf("洗牌后:%v, %d张\n", newCards, len(newCards))

	// //---发牌
	data := computer.DispactherData(newCards)
	fmt.Printf("3个人牌-底牌:%v, %d(人+底牌)\n\n", data, len(data))

	// //---发牌后排序
	// data2 := computer.CardsSortByValue(data)
	// fmt.Printf("---排序:%v, %d张\n\n", data2, len(data2))

	// //---发牌后排序
	data3 := computer.CardsSortByValue(newCards)
	fmt.Printf("-排序2-:%v, %d张\n\n", data3, len(data3))

	// utilFunc()
}

//测试
func utilFunc() {
	str := util.RandomString(5)
	fmt.Printf("-随机数-:%v, %d个\n", str, len(str))

	str2 := util.Concat(6)
	fmt.Printf("-随机数-:%v, %d个\n", str2, len(str2))
}
