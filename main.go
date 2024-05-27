// main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"math-skills/skills" // Import your local package
)

func main() {
	// reading the input data from the data.txt file
	inputStr, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("INVALID FILE", err)
		return
	}
	if len(inputStr) == 0 {
		fmt.Println("Data.txt IS EMPTY")
		return
	}
	// converting the input data form a string to a float64
	numbers := strings.Split(string(inputStr), "\n")
	var val []float64
	for _, num := range numbers {
		if num == "" {
			continue
		}
		n, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Println("INVALID INPUT IN THE FILE")
			return
		}
		val = append(val, n)
	}

	// Call the correct function from the skills package and return the result as float with 0 decimal place
	avg := skills.Avarage(val)
	fmt.Printf("Average: %0.f \n", avg)

	med := skills.Median(val)
	fmt.Printf("Median: %0.f \n", med)

	variance := skills.Variance(val)
	fmt.Printf("Variance: %0.f \n", variance)

	stdDev := skills.StandardDeviation(val)
	fmt.Printf("Standard Deviation: %0.f \n", stdDev)
}
