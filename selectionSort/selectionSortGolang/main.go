package main

import "log"

func main() {
	nums := []int{-10, 0, 1, 2, 3, 4, 6, 100, -90, -9890, 989999}
	selectionSort(nums)

	if -10001 < -10000 {
		log.Println("aaaaa")
	}

	//log.Println(selectionSort(nums))
}

func selectionSort(arr []int) {
	for i := range arr {
		smallest := i
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] < arr[smallest] { // encontrar o menor elemento entre os elementos ainda não ordenados
				// por a conparacao é com o `smallest`

				smallest = j
			}
		}

		aux := arr[i]          // Recebendo o "maior" valor
		arr[i] = arr[smallest] // indice "menor" recebendo o valor "menor" que estava alocado de forma desordenado no array
		arr[smallest] = aux    // indice "maior" recebendo o valor "maior" que estava alocado fora de ordem, estava alocado "antes" de valores menores que ele

	}

	log.Println(arr)
}
