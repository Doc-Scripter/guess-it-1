package main

import (
	"bufio"
	"fmt"
	"guess/calculations"
	"math"
	"os"
	"strconv"
	"strings"
)
//checks input data format
//calculations will only use previous 4 numbers in the input series
//minimum of 2 inputs are needed
// Calculate skewness and adjust if needed
//if skewed is false it will use normal distribution formula
//else it will use skewed distribution formula

func main() {
	if len(os.Args)!=1{
		fmt.Println("Usage: go run main.go")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	var numbers []float64
	isSkewed := false
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
			return
		}

		numbers = append(numbers, num)
		if len(numbers) > 4 {
			numbers = numbers[len(numbers)-4:]
		}
		if len(numbers) > 1 {

			skewness := calculations.Skewness(numbers)
			if math.Abs(skewness) >= 0.5 {
				isSkewed = true
			}
			
			rangeStart, rangeEnd := calculations.GuessRange(numbers, num, isSkewed)
			fmt.Printf("%d %d\n", rangeStart, rangeEnd)

		}
	}
}
