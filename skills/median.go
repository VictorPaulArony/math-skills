package skills

import "sort"

func Median(input []float64) float64 {
	// sort.Ints(input)
	// if len(input) % 2 == 0 {
	//     return (input[len(input) / 2] + input[len(input) / 2 - 1]) / 2
	// } else {
	//     return input[len(input) / 2]
	// }
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
