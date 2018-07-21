package sliceandarray

// Sum get a slice of numbers and return their sum
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll return all sum of every slice of numbers
func SumAll(numbersToSum ...[]int) (sums []int) {
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

// SumAllTails return all sum of every slice of numbers except first number
func SumAllTails(numbersToSumTails ...[]int) (sums []int) {
	for _, numbers := range numbersToSumTails {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return
}
