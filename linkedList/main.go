package main

import "fmt"

type Node struct {
	Val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(value int) {
	newNode := &Node{Val: value}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Display() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)

	list.Display()

	prev1 := searchNode(list.head, 20)
	prev2 := searchNode(prev1.next, 30)

	node1 := prev1.next
	node2 := prev2.next

	prev1.next = node2
	prev2.next = node1

	node1.next, node2.next = node2.next, node1.next

	current := list.head

	for current != nil {
		fmt.Println(current.Val)
		current = current.next
	}
	//list.Display()

}

func searchNode(node *Node, value int) *Node {

	if node == nil {
		return nil
	}

	if node.next != nil && node.next.Val == value {
		return node
	}

	return searchNode(node.next, value)

}
