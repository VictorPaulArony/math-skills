package skills

// This is the function that return variance σ2=1NN∑i=1(xi−μ)2 σ 2 = 1 N ∑ i = 1 N ( x i − μ ) 2
func Variance(input []float64) float64 {
	// var res int
	// if input == nil {
	//     return 0
	// }
	// count := 0
	// for _, v := range input {
	//     count += v
	// }
	//  res = count / len(input)
	//  return res
	var res float64
	l := float64(len(input))
	sum := 0.0
	avg := Avarage(input)
	if input == nil {
		return 0
	}
	// loop through to get 1 N ∑ i = 1 N ( x i − μ ) 2
	for _, nums := range input {
		sum += (nums - avg) * (nums - avg)
		res = sum / l
	}
	return res
}

// func main() {
// 	var input []int
// 	fmt.Println(variance(input))
// }
