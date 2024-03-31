package main

import "log"

func main() {
	//log.Println(fact(5))

	nums := []int{2, 10, 4}

	log.Println(sum(nums))

}

func fact(x int) int {
	log.Println(x)
	if x == 0 {
		return 1
	}

	return x * fact(x-1)
}

func sum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return nums[0] + sum(nums[1:])

}
