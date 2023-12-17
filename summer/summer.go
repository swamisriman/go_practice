package summer

func Sum(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
			continue
		}
		sums = append(sums, Sum(slice[1:]))
	}
	return sums
}
