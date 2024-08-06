package calculations

import (
	"testing"

	// calculations "guess/student/calculations"
)

func TestGuessRange(t *testing.T) {
	test := []struct {
		want  int
		want1 int
	}{
		{-235, 547},
		{-235, 547},
	}
	var numbers []float64 = []float64{23, 432, 12}
	var currentNum float64 = 23
	var isSkewed bool = false
	got, got1 := GuessRange(numbers, currentNum, isSkewed)
	for _, tests := range test {
		if got != tests.want || got1 != tests.want1 {
			t.Errorf("got %d %d want %d %d ",got,got1,tests.want,tests.want1)
		} 
	}
}
