package calculations

import "math"

//use formula for normal distribution
// Adjust for skewed distribution
// Handle case if the skewness condition changes 
func GuessRange(num []float64, currentNum float64, isSkewed bool) (int, int) {
	mean := Mean(num)
	standardDev := StandardDev(num)
	skewness := Skewness(num)

	var rangeStart, rangeEnd float64

	if isSkewed {
		rangeStart = mean - 2*standardDev 
		rangeEnd = mean + 2*standardDev   
	} else {
		if math.Abs(skewness) < 0.5 { 
			rangeStart = math.Max(currentNum-standardDev, mean-standardDev)
			rangeEnd = math.Min(currentNum+standardDev, mean+standardDev)
		} else {
			rangeStart = mean - 2*standardDev
			rangeEnd = mean + 2*standardDev
		}
	}

	return int(math.Round(rangeStart)), int(math.Round(rangeEnd))
}
