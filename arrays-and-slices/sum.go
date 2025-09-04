package main

func Sum(nums []int) int {
	sum := func(x, y int) int {
		return x + y
	}

	return Reduce(nums, sum, 0)
}

// the three dots is to say that this function can take a variable number of arguments
func SumAll(nums ...[]int) []int {
	sumAll := func(sum, num []int) []int {
		return append(sum, Sum(num))
	}

	return Reduce(nums, sumAll, []int{})
}

func SumAllTails(nums ...[]int) []int {
	sumTails := func(sums, num []int) []int {
		switch len(num) {
		case 0:
			sums = append(sums, 0)
		case 1:
			sums = append(sums, num[0])
		default:
			sums = append(sums, Sum(num[1:]))
		}
		return sums
	}

	return Reduce(nums, sumTails, []int{})

}

func Reduce[A any](collection []A, f func(A, A) A, intialValue A) A {
	var result = intialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
