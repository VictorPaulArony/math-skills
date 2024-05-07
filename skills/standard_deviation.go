package skills

import "math"

func StandardDeviation(input []float64) float64 {
	var res float64
	if input == nil {
		return 0
	}
	res = math.Sqrt(Variance(input))
	return res
}
