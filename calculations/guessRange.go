package calculations

import "math"

func GuessRange(num []float64, currentNum float64, isSkewed bool) (int, int) {
	mean := Mean(num)
	standardDev := StandardDev(num)
	skewness := Skewness(num)

	var rangeStart, rangeEnd float64

	if isSkewed {
		// Adjust for skewed distribution
		rangeStart = mean - 2*standardDev // Example adjustment
		rangeEnd = mean + 2*standardDev   // Example adjustment
	} else {
		if math.Abs(skewness) < 0.5 { // Assuming normal distribution
			rangeStart = math.Max(currentNum-standardDev, mean-standardDev)
			rangeEnd = math.Min(currentNum+standardDev, mean+standardDev)
		} else {
			// Handle case if the skewness condition changes (normally shouldn't happen in this design)
			rangeStart = mean - 2*standardDev
			rangeEnd = mean + 2*standardDev
		}
	}

	return int(math.Round(rangeStart)), int(math.Round(rangeEnd))
}
