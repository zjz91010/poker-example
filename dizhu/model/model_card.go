package model

/**
扑克结构体
**/
type CardData struct {
	Color int //花色1-4方梅红黑,5大小王
	Point int //牌点数A-2-3-...-10-J-Q-K-小王14-大王15
	Value int //(花色*100)+牌点数，方便带花色比较大小[Value/100=整除是花色,Value%100=取模是牌点数]
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
为了显示牌(101-113:方块A-K,201-213:梅花A-K,301-313:红桃A-K,401-413:黑桃A-K,514:小王,515:大王)
向上取整math.Ceil(),向下取整math.Floor(),取完整后返回的并不是真正的整数，而是float64 类型，所以如果需要int 类型的话需要手动转换
107/100=1(整除=花色),107%100=7(取模=牌点数)
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
