// main.go
package main

import (
	"fmt"
	"io/ioutil"

	"math-skills/skills" // Import your local package
)

func main() {
	inputStr, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("INVALID FILE", err)
		return
	}
	inputFloat64 := skills.StringToInteger(string(inputStr))
	
	// Call the correct function from the skills package
	avarage := skills.Avarage([]float64{inputFloat64})
	median := skills.Median([]float64{avarage})
	variance := skills.Variance([]float64{median})
	std := skills.StandardDeviation([]float64{variance})
	fmt.Printf("Average: %.2f\n", avarage)
	fmt.Printf("Median: %.2f\n", median)
	fmt.Printf("Variance: %.2f\n", variance)
	fmt.Printf("Standard Deviation: %.2f\n", std)
}
