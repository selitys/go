package main

func Sum(numbers []int) int {
	sum := 0
	// range lets you iterate over an array.
	// On each iteration, range returns two values - the index and the value.
	// We are choosing to ignore the index value by using _ blank identifier.
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Go can let you write variadic functions that can take a variable number of arguments.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// Safely handle empty slices
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// Slices can be sliced! The syntax is slice[low:high].
			// If you omit the value on one of the sides of the : it captures everything to that side of it
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
