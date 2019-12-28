package arrays

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(numbers ...[]int) (sums []int) {
	length := len(numbers)
	sums = make([]int, length)
	for i, arr := range numbers {
		sums[i] = Sum(arr)
	}
	return
}

func SumAllTails(numbers ...[]int) (sums []int) {
	length := len(numbers)
	sums = make([]int, length)
	for i, arr := range numbers {
		sums[i] = Sum(arr[1:])
	}
	return
}
