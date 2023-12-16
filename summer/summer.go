package summer

func Sum(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}
