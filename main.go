package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	G "guess/calculations"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var numbers []float64
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Can not read file", err)
			return
		}
		input = strings.TrimSpace(input)

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input:", err)
			return
		}
		numbers = append(numbers, num)
		
		
		
			rangeStart, rangeEnd := G.GuessRange(numbers, num)
			fmt.Printf("%d %d  \n", rangeStart, rangeEnd)
		
	}
}
