package shared

func MinInt(nums *[]int) (int, int) {
	min := (*nums)[0]
	pos := 0

	for i, num := range *nums {
		if num < min {
			min = num
			pos = i
		}
	}

	return pos, min
}

func MaxInt(nums *[]int) (int, int) {
	max := (*nums)[0]
	pos := 0

	for i, num := range *nums {
		if num > max {
			max = num
			pos = i
		}
	}

	return pos, max
}
