/* 2. Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout. */

package main

import (
	"fmt"
	"sync"
)

//Реализуем вариант с функцией (будем запускать её в горутине), пишущей квадраты в канал
//и закрывающей его по окончанию записи (элементов массива)
//в качестве параметра принимает chan<- - указание, что канал только для записи
func getPowChan(c chan<- int, in []int) {
	for _, v := range in {
		c <- v * v
	}
	close(c)
}

//Реализуем вариант с функцией, запускающей возведение в квадрат в отдельной горутине (предположим это был бы трудоемкий процесс)
//для каждого элемента. Полученное значение горутина записывает в соответствующую ячейку (элемент) слайса.
//Так как мы гарантируем, что каждая горутина работает с отдельной ячейкой слайса (ячейкой памяти) - это безопасно.
//Ожидаем выполнение всех горутин при помощи WaitGroup и возвращаем заполненный слайс.
func getPowWG(in []int) []int {
	pows := make([]int, len(in))
	var wg sync.WaitGroup

	for index, value := range in {
		wg.Add(1)
		go func(value int, res []int, index int) {
			defer wg.Done()
			pows[index] = value * value
		}(value, pows, index)
	}
	wg.Wait()
	return pows
}

func main() {
	//Создадим слайс, сразу проиницилизировав его значениям, так как он задан строго по заданию
	inputArr := []int{2, 4, 6, 8, 10}

	res := getPowWG(inputArr)
	fmt.Println(res)

	channel := make(chan int)
	go getPowChan(channel, inputArr)

	for pow := range channel {
		fmt.Println(pow) // используем fmt.Println - стандартно выводит в stdout
	}
}
