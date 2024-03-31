package binary

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(list, 7))
}

// Busca em uma lista ORDENADA
func binarySearch(list []int, elem int) int {
	inicio := 0
	fim := len(list) - 1

	meio := 0

	//Percorre a lista
	for inicio <= fim {

		// Encontra o meio da lista e usa como base para o ˜palpide˜
		meio = (inicio + fim) / 2

		if list[meio] == elem {
			return meio

		} else if list[meio] < elem {
			// Se o elemento for maior que o palpite, soma 1 no inicio, ou seja o elemento está na metade da direta do array
			inicio = meio + 1

		} else if list[meio] > elem {
			// Se o elemento for menor que o palpite, diminui 1 no fim, ou seja  o elemento está na metade da esquerda do array
			fim = meio - 1
		}
	}

	return meio
}
