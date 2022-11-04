package computer

import (
	"doudizhutest/enum"
	"doudizhutest/model"
	"doudizhutest/util"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

/**
一、斗地主牌面分析
斗地主需要54张牌。

2-10 黑桃、红桃、梅花、方片各4张。
J、Q、K、A 黑桃、红桃、梅花、方片各4张。
大小王各1张。
在斗地主中，牌的花色不影响。所以，在牌面比对时，不需要单独比对花色。而单张牌面值的大小顺序为: 大王>小王>2>A>K>Q>J>10……3
鉴于此，牌面的表达可以用以下方式来规定：
A：黑桃 B：红桃 C：梅花 D：方片

扑克原始值	映射值
3-10	3-10数字
J		11
Q		12
K		13
A		14
2		15
小王		Q16
大王		K17
例如：
A14----->黑桃A
C9----->梅花9
二、游戏初始化拆分成3大块

1：构造一副牌，2：洗牌，3：发牌
**/

//定义一个常量方便参数修改
const (
	COLOR_A = "A"
	COLOR_B = "B"
	COLOR_C = "C"
	COLOR_D = "D"
	COLOR_Q = "Q16" //小王
	COLOR_K = "K17" //大王
)

/**
构造一副牌
**/
func CreateNew() []string {
	numbers := make([]string, 54) //构造一个大小为54的数组
	start := 0
	for i := 3; i <= 16; i++ {
		if i == 16 {
			//i为16说明已经到达大小王
			numbers[start] = COLOR_Q
			numbers[start+1] = COLOR_K //直接构造大小王
		} else {
			numbers[start] = COLOR_A + strconv.Itoa(i)
			numbers[start+1] = COLOR_B + strconv.Itoa(i)
			numbers[start+2] = COLOR_C + strconv.Itoa(i)
			numbers[start+3] = COLOR_D + strconv.Itoa(i)
			start += 4 //每造一套单值牌，游标移4位
		}
	}
	return numbers
}

/**
洗牌
洗牌就是将牌原有的顺序打乱，形成新的顺序的牌。主要利用随机数来处理
**/
func Shuffle(vals []string) {
	r := rand.New(rand.NewSource(time.Now().Unix())) //根据系统时间戳初始化Random
	for len(vals) > 0 {
		n := len(vals)                                          //根据卡牌数组长度遍历
		randIndex := r.Intn(n)                                  //得到随机数index
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1] //最后一张牌和第randIndex张牌互换
		vals = vals[:n-1]
	}
}

/**
发牌
发牌可以说是斗地主开始前的最后一个环节(不包含叫地主抢地主)，发牌是要将牌先均分给3个玩家(保留3张底牌)，并从玩家中随机抽取一位玩家为地主。
首先，将牌分成4部分:
玩家一：17张牌
玩家二：17张牌
玩家三：17张牌
底牌：3张
**/

/**
发牌order为玩家位置顺序
order==0 玩家1次序
order==1 玩家2次序
order==2 玩家3次序
order==3 底牌次序
**/
func Dispacther(order int, vals []string) []string {
	var playCards []string
	//判断玩家次序是否正确(3个人玩)
	if order < 0 || order > 3 {
		return []string{}
	} else {
		size := 17 //默认总长度
		if order == 3 {
			size = 3 //最后3张是低牌
		}
		for i := 0; i < size; i++ {
			//洗牌后根据顺序发牌
			playCards = append(playCards, vals[order*17+i])
		}
	}
	return playCards
}

/**
根据牌面数量判断牌面类型
**/
func ParseCardsInSize(plays []string) model.CardShow {
	cardShow := model.CardShow{
		ShowValue: plays, //出牌值
		ShowTime:  util.GetNowTime(),
	}
	playLen := len(plays)
	if playLen > 0 {
		if playLen > 2 {
			ParseCardsType(plays, &cardShow) //判断出啥牌
		} else {
			switch playLen {
			case 1: //随便出一张
				cardShow.CardTypeStatus = enum.SINGLE
				cardShow.CompareValue = GetCardValue(plays[0])
				cardShow.MaxCount = 1
				cardShow.MaxValues = []int{cardShow.CompareValue}

			case 2: //随便出2张
				//先判断是否是王炸
				if plays[0] == COLOR_Q && plays[1] == COLOR_K {
					cardShow.CardTypeStatus = enum.KING_BOMB
					cardShow.CompareValue = GetCardValue(plays[0])
					cardShow.MaxCount = 2
					cardShow.MaxValues = []int{cardShow.CompareValue}
				} else {
					//否则是对子，或出错
					ParseCardsType(plays, &cardShow) //判断出啥牌
				}
			}
		}
	} else {
		cardShow.CardTypeStatus = enum.ERROR_TYPE
	}

	return cardShow
}

/**
获取牌面类型
**/
func ParseCardsType(cards []string, cardShow *model.CardShow) {
	mapCard, maxCount, maxValues := ComputerValueTimes(cards)
	cardShow.MaxCount = maxCount
	cardShow.MaxValues = maxValues
	cardShow.CardMap = mapCard
	cardShow.CompareValue = maxValues[len(maxValues)-1]
	switch maxCount {
	case 4:
		if maxCount == len(cards) {
			cardShow.CardTypeStatus = enum.KING_BOMB
			fmt.Println("炸弹")
		} else if len(cards) == 6 {
			cardShow.CardTypeStatus = enum.FOUR_TWO
			fmt.Println("四带二")
		} else {
			cardShow.CardTypeStatus = enum.ERROR_TYPE
			fmt.Println("不合法出牌")
		}
		break
	case 3:
		alive := len(cards) - len(maxValues)*maxCount
		if len(maxValues) == alive {
			if len(maxValues) == 1 {
				cardShow.CardTypeStatus = enum.THREE_AND_ONE
				fmt.Println("三带一")
			} else if len(maxValues) > 1 {
				if IsContinuity(mapCard, false) {
					cardShow.CardTypeStatus = enum.PLANE
					fmt.Printf("飞机%d", len(maxValues))
				} else {
					cardShow.CardTypeStatus = enum.ERROR_TYPE
					fmt.Println("非法飞机")
				}
			}
		} else if alive == 0 {
			if len(maxValues) > 1 {
				if IsContinuity(mapCard, false) {
					cardShow.CardTypeStatus = enum.PLANE_EMPTY
					fmt.Printf("三不带飞机%d", len(maxValues))
				} else {
					cardShow.CardTypeStatus = enum.ERROR_TYPE
					fmt.Println("非法三不带飞机")
				}

			} else {
				cardShow.CardTypeStatus = enum.THREE
				fmt.Println("三不带")
			}
		} else {
			cardShow.CardTypeStatus = enum.ERROR_TYPE
			fmt.Println("不合法飞机或三带一")
		}
		break
	case 2:
		if len(maxValues) == (len(cards) / 2) {
			if len(maxValues) > 1 {
				if IsContinuity(mapCard, false) && len(maxValues) > 2 {
					cardShow.CardTypeStatus = enum.DOUBLE_ALONE
					fmt.Printf("%d连队", len(maxValues))
				} else {
					cardShow.CardTypeStatus = enum.ERROR_TYPE
					fmt.Println("非法连对")
				}
			} else if len(maxValues) == 1 {
				cardShow.CardTypeStatus = enum.DOUBLE
				fmt.Printf("对%d", GetCardValue(cards[0]))
			}
		} else {
			cardShow.CardTypeStatus = enum.ERROR_TYPE
			fmt.Println("不合法出牌")
		}
		break
	case 1:
		if IsContinuity(mapCard, true) && len(cards) >= 5 {
			cardShow.CardTypeStatus = enum.SINGLE_ALONE
			fmt.Printf("%d顺子", len(mapCard))
		} else {
			cardShow.CardTypeStatus = enum.ERROR_TYPE
			fmt.Println("非法顺子")
		}
		break
	}
}

/**
* 获取顺序的key值数组
 */
func GetOrderKeys(cardMap map[int]int, isSingle bool) []int {
	var keys []int
	for key, value := range cardMap {
		if (!isSingle && value > 1) || isSingle {
			keys = append(keys, key)
		}
	}
	sort.Ints(keys)
	return keys
}

/**
* 计算牌面值是否连续
 */
func IsContinuity(cardMap map[int]int, isSingle bool) bool {
	keys := GetOrderKeys(cardMap, isSingle)
	lastKey := 0
	for i := 0; i < len(keys); i++ {
		if (lastKey > 0 && (keys[i]-lastKey) != 1) || keys[i] == 15 {
			return false
		}
		lastKey = keys[i]
	}
	if lastKey > 0 {
		return true
	} else {
		return false
	}
}

/**
* 计算每张牌面出现的次数
* mapCard 标记结果
* MaxCount 出现最多的次数
* MaxValues 出现次数最多的所有值
 */
func ComputerValueTimes(cards []string) (mapCard map[int]int, MaxCount int, MaxValues []int) {
	newMap := make(map[int]int)
	if len(cards) == 0 {
		return newMap, 0, nil
	}
	for _, value := range cards {
		cardValue := GetCardValue(value)
		if newMap[cardValue] != 0 {
			newMap[cardValue]++
		} else {
			newMap[cardValue] = 1
		}
	}
	var allCount []int //所有的次数
	var maxCount int   //出现最多的次数
	for _, value := range newMap {
		allCount = append(allCount, value)
	}
	maxCount = allCount[0]
	for i := 0; i < len(allCount); i++ {
		if maxCount < allCount[i] {
			maxCount = allCount[i]
		}
	}
	var maxValue []int
	for key, value := range newMap {
		if value == maxCount {
			maxValue = append(maxValue, key)
		}
	}
	sort.Ints(maxValue)
	return newMap, maxCount, maxValue
}

/**
获取牌面值
**/
func GetCardValue(card string) int {
	stringValue := util.Substring(card, 1, len(card)) //截取字符串数字
	value, err := strconv.Atoi(stringValue)           //字符串转int
	if err == nil {
		return value
	}
	return -1
}
