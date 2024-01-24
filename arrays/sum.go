package array

func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func SumAll(numbersToSum ...[]int) []int {
	totals := []int{}

	for _, numbers := range numbersToSum {
		totals = append(totals, Sum(numbers))
	}

	return totals
}

func SumAllTails(numbersToSum ...[]int) []int {
	totals := []int{}

	for _, numbers := range numbersToSum {
		if len(numbers) != 0 {
			totals = append(totals, Sum(numbers[1:]))
		} else {
			totals = append(totals, 0)
		}
	}

	return totals
}
