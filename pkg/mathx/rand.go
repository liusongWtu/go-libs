package mathx

import (
	"math/rand"
	"time"
)

// RandomFloat 随机[min,max],随机小数位数decimalPlaces
func RandomFloat(min, max float64, decimalPlaces int) float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := min + rand.Float64()*(max-min)

	shift := float64(1)
	for i := 0; i < decimalPlaces; i++ {
		shift *= 10
	}
	return float64(int(randomNumber*shift)) / shift
}
