package model

/**
扑克结构体
**/
type CardData struct {
	Color int //花色1-4方梅红黑,5大小王
	Point int //牌值点数A-2-3-...-10-J-Q-K-小王14-大王15
	Value int //(花色*100)+牌值点数，方便比较大小
}

/**
构造函数，传入54的牌值
**/
func NewCardData(index int) *CardData {
	var cd CardData
	cd.Color = ((index - 1) / 13) + 1
	cd.Point = toCardPoint(index)
	cd.Value = toCardValue(index)
	return &cd
}

/**
传入牌的序号1-54
为了显示牌101-115,(101:方块A,201:梅花A,301:红桃A,401:黑桃A,514:小王,515:大王)
**/
func toCardValue(index int) int {
	val := 0
	if index <= 0 {
		return val
	}
	val = (((index-1)/13)+1)*100 + (((index - 1) % 13) + 1)
	return val
}

/**
传入牌的序号1-54
为了显示牌1-15,["1","2","3",,,"10","J","Q","K","小王","大王"]
**/
func toCardPoint(index int) int {
	val := 0
	if index <= 0 {
		return val
	}
	switch index {
	case 54:
		val = 15 //大王
	case 53:
		val = 14 //小王
	default:
		val = ((index - 1) % 13) + 1
	}
	return val
}

/**
创造一副牌
num为传入是52或54（去除大小王52张 或 包括大小王54张）
**/
func CreateNewCards(num int) []CardData {
	var cards []CardData
	if num <= 0 {
		return cards
	}
	for i := 1; i <= num; i++ {
		cards = append(cards, *NewCardData(i))
	}
	return cards
}
