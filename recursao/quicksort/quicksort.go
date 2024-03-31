package main

import "fmt"

func main() {
	fmt.Println(quicksort([]int{10, 5, 2, 3}))
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]

	//pivot := nums[middle]
	left := []int{}
	right := []int{}

	for _, value := range nums[1:] {

		if value <= pivot {
			left = append(left, value)
			continue
		}

		right = append(right, value)
	}

	fmt.Println(left)
	fmt.Println(right)

	var result []int

	result = append(result, quicksort(left)...)
	result = append(result, pivot)
	result = append(result, quicksort(right)...)

	return result

}
