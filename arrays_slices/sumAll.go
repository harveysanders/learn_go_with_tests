package main

func SumAll(numsToSum ...[]int) []int {
	sums := make([]int, 0)
	for _, v := range numsToSum {
		sums = append(sums, Sum(v))
	}
	return sums
}
