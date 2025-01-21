package arrays

func Sum(numbers []int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	lengthOfNumbers := len(numbersToSum)
	result := make([]int, lengthOfNumbers)

	for i, array := range numbersToSum {
		result[i] = Sum(array)
	}
	return result
}

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		tail := nums[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}