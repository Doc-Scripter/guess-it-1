package calculations

import "math"

func GuessRange(num []float64, currentNum float64, nextNum float64) (int, int) {
	mean := Mean(num)
	standardDev := StandardDev(num)

	rangeStart := math.Max(currentNum-standardDev, mean-standardDev)
	rangeEnd := math.Min(nextNum+standardDev, mean+standardDev)

	rangeStart = math.Max(rangeStart, currentNum)
	rangeEnd = math.Min(rangeEnd, nextNum)

	if rangeStart > rangeEnd {
		rangeStart, rangeEnd = rangeEnd, rangeStart
	}
	return int(math.Round(rangeStart)), int(math.Round(rangeEnd))
}
