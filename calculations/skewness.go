package calculations

import "math"

func Skewness(data []float64) float64 {
	mean := Mean(data)
	stdDev := StandardDev(data)
	n := float64(len(data))
	sum := 0.0
	for _, x := range data {
		sum += math.Pow(x-mean, 3)
	}
	return (n * sum) / ((n - 1) * (n - 2) * math.Pow(stdDev, 3))
}
