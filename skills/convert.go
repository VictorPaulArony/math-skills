package skills

import (
	"strconv"
)
// Function to convert string to float64
func StringToInteger(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
    if err!= nil {
        return 0
    }
    return res
}
