package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node3 := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	head := ListNode{Val: 1, Next: &node1}

	fmt.Println(isPalindrome(&head))

}

func isPalindrome(head *ListNode) bool {
	//fmt.Println(getArray(head))

	values := getArray(head)
	left := 0
	right := len(values) - 1

	for left <= right {

		if values[left] != values[right] {
			return false
		}

		left++
		right--

	}

	return true
}

func getArray(head *ListNode) []int {

	element := head
	result := []int{}

	for element != nil {
		result = append(result, element.Val)
		element = element.Next
	}

	return result
}
