package shared

func AbsInt(number int) int {
	if number < 0 {
		return -1 * number
	}

	return number
}

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

func SumInt(nums *[]int) (sum int) {
	for _, v := range *nums {
		sum += v
	}

	return
}

func SumVerticalInt(nums *[][]int, col int, ch chan int) {
	sum := 0

	for i := range *nums {
		sum += (*nums)[i][col]
	}

	ch <- sum
}
