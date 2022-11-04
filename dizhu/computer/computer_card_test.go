package computer

import (
	"dizhu/computer"
	"dizhu/model"
	"fmt"
	"log"
	"testing"
)

//黑桃♠,红桃♥,梅花♣,方块♦,小王🃏,大王🃏
var CardColor []string = []string{"", "♦", "♣", "♥", "♠", "🃏"}
var CardValue []string = []string{"", "A", "J", "Q", "K", "小王", "大王"}

func TestCardsShow(t *testing.T) {
	cards := CreateNewCards()
	log.Printf("构造一副牌:\n%v\n, %T,%d张\n", cards, cards, len(cards))

	cards2 := computer.CardsSortByValue(cards)
	log.Printf("大小花色排序:\n%v\n, %T,%d张\n", cards2, cards2, len(cards2))

	showCards := ConcatCardsVal(cards)
	t.Logf("花色:%v \n %d张\n", showCards, len(showCards))
}

/**
拼接牌点数方便显示🃏-♠1-...-♠13
**/
func ConcatCardsVal(vals []model.CardData) []model.CardData {
	showCard := []string{}

	for i := range vals {
		flower := ""
		switch vals[i].Value {
		case 1:
			flower = fmt.Sprintf("%v%v", CardColor[int(vals[i].Color)], CardValue[1])
		case 11, 12, 13, 14, 15:
			flower = fmt.Sprintf("%v%v", CardColor[int(vals[i].Color)], CardValue[vals[i].Value-9])
		default:
			flower = fmt.Sprintf("%v%v", CardColor[int(vals[i].Color)], vals[i].Value)
		}
		showCard = append(showCard, flower)

		// log.Printf("i=%d:%v \n", i+1, showCard)

	}
	log.Printf("showCard:%v \n %v \n", showCard, len(showCard))

	return vals
}
