package main

import "fmt"

func main() {
	// Teste 1: nums = [4,5,6,7,0,1,2], target = 0
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	target1 := 0
	expected1 := 4
	result1 := search(nums1, target1)
	if result1 != expected1 {
		fmt.Printf("Para nums=%v, target=%d, esperado=%d, obtido=%d", nums1, target1, expected1, result1)
	}

	// Teste 2: nums = [4,5,6,7,0,1,2], target = 3 (elemento não presente)
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	expected2 := -1
	result2 := search(nums2, target2)
	if result2 != expected2 {
		fmt.Printf("Para nums=%v, target=%d, esperado=%d, obtido=%d", nums2, target2, expected2, result2)
	}

	// Teste 3: nums = [1], target = 0 (array com um elemento)
	nums3 := []int{1}
	target3 := 0
	expected3 := -1
	result3 := search(nums3, target3)
	if result3 != expected3 {
		fmt.Printf("Para nums=%v, target=%d, esperado=%d, obtido=%d", nums3, target3, expected3, result3)
	}

	// Teste 4: nums = [4,5,6,7,8,1,2,3], target = 8 (último elemento)
	nums4 := []int{4, 5, 6, 7, 8, 1, 2, 3}
	target4 := 8
	expected4 := 4
	result4 := search(nums4, target4)
	if result4 != expected4 {
		fmt.Printf("Para nums=%v, target=%d, esperado=%d, obtido=%d", nums4, target4, expected4, result4)
	}

	// Teste 5: nums = [5,6,7,8,1,2,3,4], target = 1 (primeiro elemento)
	nums5 := []int{5, 6, 7, 8, 1, 2, 3, 4}
	target5 := 1
	expected5 := 4
	result5 := search(nums5, target5)
	if result5 != expected5 {
		fmt.Printf("Para nums=%v, target=%d, esperado=%d, obtido=%d", nums5, target5, expected5, result5)
	}
}

func search(nums []int, target int) int {
	inicio := 0
	fim := len(nums) - 1

	for inicio <= fim {

		meio := (inicio + fim) / 2

		if nums[meio] == target {
			return meio

		} else if nums[meio] >= nums[inicio] { // target do lado esquerdo
			
			if nums[inicio] <= target && target < nums[meio] {
				fim = meio - 1
			} else {
				inicio = meio + 1
			}

		} else {
			if nums[meio] < target && target <= nums[fim] { // target do lado direito
				inicio = meio + 1

			} else {
				fim = meio - 1
			}
		}
	}

	return -1
}
