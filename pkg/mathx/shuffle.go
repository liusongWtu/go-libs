package mathx

import (
	"fmt"
	"math/rand"
	"time"
)

func demo() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{1, 2, 3, 4, 5} //也可以是数组
	fmt.Println(arr)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)
}

func Shuffle(weights []float64) []int {
	count := len(weights)
	if count == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())

	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = i
	}

	for i := 0; i < count-1; i++ {
		totalWeight := 0.0
		for j := i; j < count; j++ {
			totalWeight += weights[j]
		}

		random := rand.Float64() * totalWeight
		accumulator := 0.0
		selected := 0
		for k := i; k < count; k++ {
			accumulator += weights[k]
			if random < accumulator {
				selected = k
				break
			}
		}

		weights[i], weights[selected] = weights[selected], weights[i]
		result[i], result[selected] = result[selected], result[i]
	}
	return result
}
