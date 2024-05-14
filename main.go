// main.go
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"math-skills/skills" // Import your local package
)

func main() {
	inputStr, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("INVALID FILE", err)
		return
	}
	if len(inputStr) == 0 {
        fmt.Println("Data.txt IS EMPTY")
		return
    }

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

	// Call the correct function from the skills package
	avg := skills.Avarage(val)
	// if avg > 184467440737095516 {
	// 	fmt.Println("TO MANY VALUES TO RETURN")

	// }
	fmt.Printf("Average: %0.f \n", avg) // math.Round(avg))
	med := skills.Median(val)
	// if med > 184467440737095516 {
	// 	fmt.Println("TO MANY VALUES TO RETURN")

	// }
	fmt.Printf("Median: %0.f \n", med)
	variance := skills.Variance(val)
	// if variance > 184467440737095516 {
	// 	fmt.Println("TO MANY VALUES TO RETURN")
	// }
	fmt.Printf("Variance: %0.f \n", variance)
	stdDev := skills.StandardDeviation(val)
	// if stdDev > 184467440737095516 {
	// 	fmt.Println("TO MANY VALUES TO RETURN")
	// }
	fmt.Printf("Standard Deviation: %0.f \n", stdDev)
}
