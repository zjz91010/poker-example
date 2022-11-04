package computer

import (
	"dizhu/model"
	"math/rand"
	"sort"
	"time"
)

/**
获取一副新牌
**/
func CreateNewCards() []model.CardData {
	cardData := model.CreateNewCards(54)
	return cardData
}

/**
洗牌
*/
func Shuffle(vals []model.CardData) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)         //数组中有多少个元素
		randIndex := r.Intn(n) //随机生成一个key
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}

}

/**
斗地主发牌order为玩家位置顺序
order==0 玩家1次序
order==1 玩家2次序
order==2 玩家3次序
order==3 底牌次序
**/
func Dispacther(order int, vals []model.CardData) []model.CardData {
	var playCards []model.CardData
	//判断玩家次序是否正确(3个人玩)
	if order < 0 || order > 3 {
		return []model.CardData{}
	} else {
		size := 17 //默认总长度
		if order == 3 {
			size = 3 //最后3张是低牌
		}
		for i := 0; i < size; i++ {
			//洗牌后根据顺序发牌
			playCards = append(playCards, vals[order*17+i])

			//下面是每个人的牌直接排好序
			// CardsSortByValue(playCards)
		}
	}
	return playCards
}

/**
发牌（把所有玩家牌+底牌数据放在一起【未排序】）
**/
func DispactherData(vals []model.CardData) [][]model.CardData {
	var playersArr [][]model.CardData
	for i := 0; i <= 3; i++ {
		player := Dispacther(i, vals)
		playersArr = append(playersArr, player)
	}
	return playersArr
}

/**
斗地主发牌后牌排序
**/
func DispactherSortByColor(vals []model.CardData) []model.CardData {

	player := Dispacther(0, vals)

	sort.Slice(player, func(i, j int) bool {
		if player[i].Value == player[j].Value {
			return player[i].Color < player[j].Color
		}
		return player[i].Value > player[j].Value
	})

	return player
}

/**
一维
牌值点数+花色排序
**/
func CardsSortByValue(vals []model.CardData) []model.CardData {

	sort.Slice(vals, func(i, j int) bool {
		return vals[i].Value > vals[j].Value
	})
	return vals
}
