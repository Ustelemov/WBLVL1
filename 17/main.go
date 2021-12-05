/*Реализовать бинарный поиск встроенными методами языка. */
package main

import "fmt"

//number - искомый элемент в in, выходное значение - индекс элемента в слайсе, -1 - элемента нет
//in - должен быть отсортирован!
func binarySearch(in []int, searcing int) (result int) {
	middle := len(in) / 2 // индекс центрального элемента

	switch {
	case len(in) == 0:
		result = -1
	case in[middle] > searcing:
		result = binarySearch(in[:middle], searcing)
	case in[middle] < searcing:
		result = binarySearch(in[middle+1:], searcing)
		if result != -1 { //если элемент есть
			result += middle + 1 //учтем, что 0 элемент - это middle+1 элемент
		}
	default: // in[middle] == searching
		result = middle
	}
	return
}

func main() {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Input: %v\n", in)
	fmt.Printf("Index for %d is %d\n", 5, binarySearch(in, 5))
}
