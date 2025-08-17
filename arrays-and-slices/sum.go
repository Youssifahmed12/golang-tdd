package main

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// the three dots is to say that this function can take a variable number of arguments
func SumAll(nums ...[]int) []int {
	var sums []int
	for _, num := range nums {
		sums = append(sums, Sum(num))
	}
	return sums
}

func SumAllTails(nums ...[]int) []int {
	var sums []int
	for _, n := range nums {
		sums = append(sums, Sum(n[1:]))
	}
	return sums
}
