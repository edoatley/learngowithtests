package arrays

func Sum(numbers []int) (sum int) {
	
	// range lets you iterate over an array. On each iteration, range returns two values - the index and the value. 
	// We are choosing to ignore the index value by using _ blank identifier.
	// we could have used a for loop to iterate over the array, but using range is a lot cleaner.
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers) // make allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// Slices can be sliced! The syntax is slice[low:high]. If you omit the value on one of the sides of the : it captures everything to that side of it. In our case, we are saying "take from 1 to the end" with numbers[1:]
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}