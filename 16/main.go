/*Реализовать быструю сортировку массива (quicksort) встроенными методамиязыка. */
package main

import (
	"fmt"
	"math/rand"
)

func quickSort(in []int) []int {
	if len(in) < 2 { //если элементов нет или он один - сортировать не нужно
		return in
	}

	left, right := 0, len(in)-1

	pivot := rand.Int() % len(in) //возьмем случайный элемент в качестве опорного

	in[pivot], in[right] = in[right], in[pivot] //установим опорный элемент вправо

	// сделаем так чтобы слева стояли элементы меньше опорного
	for i := range in {
		if in[i] < in[right] {
			in[left], in[i] = in[i], in[left]
			left++ //в этой позиции элемент уже меньше опорного сдвигаем дальше
		}
	}

	//left - позиция последнего элемента, который больше опорного, установим в неё опорный элемент
	in[left], in[right] = in[right], in[left]

	//повторим все операции для элементов до и после опорного
	quickSort(in[:left])
	quickSort(in[left+1:])

	return in
}

func main() {
	in := []int{2, 10, 5, 101, 231, 0, 1, 23, 4, -5}
	fmt.Printf("Input: %v\n", in)
	fmt.Printf("Sorted output: %v\n", quickSort(in))
}
