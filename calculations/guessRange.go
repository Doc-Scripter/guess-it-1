package calculations

import "math"

func GuessRange(num []float64, currentNum float64) (int, int) {
	mean := Mean(num)
	standardDev := StandardDev(num)

	rangeStart := math.Max(currentNum-standardDev, mean-standardDev)
	rangeEnd := math.Min(currentNum+standardDev, mean+standardDev)

	
	return int(math.Round(rangeStart)), int(math.Round(rangeEnd))
}
