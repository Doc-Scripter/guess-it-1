package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"guess/calculations"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var numbers []float64
	isSkewed:=false
	for {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Cannot read input:", err)
			return
		}
		input = strings.TrimSpace(input)

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		numbers = append(numbers, num)

		// Calculate skewness and adjust if needed
		skewness := calculations.Skewness(numbers)
		if math.Abs(skewness) >= 0.5 {
			isSkewed = true
		}

		// Call GuessRange to get the predicted range
		rangeStart, rangeEnd := calculations.GuessRange(numbers, num, isSkewed)
		fmt.Printf("%d - %d\n", rangeStart, rangeEnd)
	}
}
