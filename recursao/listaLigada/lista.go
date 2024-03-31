package main

import (
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Criando o primeiro nó da lista
	primeiroNo := &ListNode{Val: 1}

	// Criando o segundo nó da lista
	segundoNo := &ListNode{Val: 2}

	// Ligando o primeiro nó ao segundo nó
	primeiroNo.Next = segundoNo

	// Criando o terceiro nó da lista
	terceiroNo := &ListNode{Val: 3}

	// Ligando o segundo nó ao terceiro nó
	segundoNo.Next = terceiroNo

	reverseList(primeiroNo)

}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //Quando chegar no ultimo elemento, devolve ele
		return head
	}

	result := reverseList(head.Next) // vai entrando em recursão até chegar no ultimo

	log.Println(fmt.Sprintf("Head value: %d, Result value: %d", head.Val, result.Val))
	head.Next.Next = head // adicionando o head ao next do "ultimo elemento", no exemplo: adicionando o '2' como NEXT do '3'
	head.Next = nil       // zerando o 'NEXT' do head

	return result
}
