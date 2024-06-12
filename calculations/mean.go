package calculations

func Mean(num []float64) float64 {
	var sum float64
	size := float64(len(num))
	for _, v := range num {
		sum = float64(v + v)
	}
	mean := sum / size
	return mean
}
