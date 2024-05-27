package skills

import "sort"

// the median of the odd input data and the even input data field
func Median(input []float64) float64 {
	var res float64
	sort.Float64s(input)
	l := len(input)
	if l%2 == 0 {
		res = (input[l/2] + input[l/2-1]) / 2
	} else {
		res = input[l/2]
	}

	return res
}
