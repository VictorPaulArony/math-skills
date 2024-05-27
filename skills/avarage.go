package skills

// function to calculate the avarage of the given numbers in the input list
func Avarage(input []float64) float64 {
	var res float64
	if input == nil {
		return 0.0
	}
	count := 0.0
	for _, v := range input {
		count += v
	}
	res = count / float64(len(input))

	return res
}
