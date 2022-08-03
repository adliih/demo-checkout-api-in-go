package main

import (
	"fmt"
	"sort"
)

// Given array: [10, 15, 11, 3, 2, 1, 30, 45]
// Expected result: [45,1], [30, 2], ...

func tuppling(input []int) [][]int  {
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j]; })

	n := len(input)
	var result [][]int
	for i := 0; i < n / 2; i++ {
		result = append(result, []int{ input[n - i - 1], input[i] })
	}

	return result
}

func main()  {
	fmt.Println(tuppling([]int{10, 15, 11, 3, 2, 1, 30, 45}))
}