// main.go
package main

import (
    "fmt"
	"strings"
    "io/ioutil"
    "math"
    "strconv"

    "math-skills/skills" // Import your local package
)

func main() {
    inputStr, err := ioutil.ReadFile("data.txt")
    if err != nil {
        fmt.Println("INVALID FILE", err)
        return
    }

    numbers:= strings.Split(string(inputStr), "\n")
	var val []float64
	for _, num := range numbers{
		if num == "" {
            continue
        }
    n, err := strconv.ParseFloat(num, 64)
    if err != nil {
        fmt.Println("INVALID NUMBERS", err)
        return
    }
	val = append(val, n)
}

    // Call the correct function from the skills package
    avg := skills.Avarage(val)
    fmt.Println("Average:", math.Round(avg))
    med := skills.Median(val)
    fmt.Println("Median:", math.Round(med))
    variance := skills.Variance(val)
    fmt.Println("Variance:", variance)
    stdDev := skills.StandardDeviation(val)
    fmt.Println("Standard Deviation:", stdDev)
}