package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	fmt.Println(search(nums, target))
}

func search(arr []int, target int) int {
	fmt.Println(arr)
	if len(arr) == 0 {
		return -1
	}

	middle := (len(arr) - 1) / 2

	if arr[middle] == target {
		return middle
	}

	if arr[middle] < target {
		return search(arr[middle+1:], target)
	}

	if arr[middle] > target {
		return search(arr[:middle-1], target)
	}

	return -1

}
