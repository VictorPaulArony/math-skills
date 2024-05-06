// main.go
package main

import (
	"fmt"

	"math-skills/skills" // Import your local package
)

func main() {
	skills.MyFunction() // Use the function from your local package
	fmt.Println("Hello from my main program!")
}
