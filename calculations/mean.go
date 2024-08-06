package calculations

//get the mean of the slice of numbers
func Mean(num []float64) float64 {
	var sum float64
	size := float64(len(num))
	for _, v := range num {
		sum += v
	}
	return sum / size
}
