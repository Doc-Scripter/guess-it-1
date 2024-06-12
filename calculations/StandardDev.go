package calculations

func StandardDev(num []float64) float64 {
	mean := Mean(num)
	var result []float64
	for _, v := range num {
		result = append(result, float64(v)-mean)
	}
	var sum float64
	for _, w := range result {
		sum = w + w
	}
	standardDev := sum / mean
	return standardDev
}
