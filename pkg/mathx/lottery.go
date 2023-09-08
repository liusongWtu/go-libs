package mathx

import (
	"github.com/gogf/gf/util/grand"
	"math"
)

// 传入权重数组，返回抽中的序号
func Lottery(weights []int) int {
	total := 0
	for _, weight := range weights {
		total += weight
	}

	i := grand.N(1, total)
	for index, weight := range weights {
		if i <= weight {
			return index
		}
		i -= weight
	}
	//永远不会返回-1
	return -1
}

// LotteryByRate 根据概率进行抽奖
// rate 中奖概率
// precision 中奖概率的精确度（精确到的小数位数）
func LotteryByRate(rate, precision float64) bool {
	maxF := math.Pow(10, precision)
	max := int(maxF)
	randVal := grand.Intn(max)
	rateVal := int(rate * maxF)

	return randVal <= rateVal
}
