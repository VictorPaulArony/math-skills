package skills

import (
	"strconv"
)
// Function to convert string to float64
func StringToInteger(input string) float64 {
	res, err := strconv.ParseFloat(string(input), 64)
    if err!= nil {
        return 0
    }
    
    return res
}
