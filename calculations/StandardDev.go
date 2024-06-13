package calculations

import (
	"math"
)

func StandardDev(num []float64) float64 {
	mean := Mean(num)
	var result []float64
	for _, v := range num {
		result = append(result, v-mean)
	}
	var sum float64
	for _, w := range result {
		sum += w * w
	}
	variance := sum / float64(len(num))
	return math.Sqrt(variance)
}
