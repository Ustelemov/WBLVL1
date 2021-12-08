/*Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой. */
package main

import (
	"fmt"
	"strings"

	"github.com/Ustelemov/WBLVL1/12/set"
)

func checkUnique(in string) bool {
	//Приведем все символы к нижнему регистру, функция должна быть регистронезависимой
	lower := strings.ToLower(in)

	//Создаем реализацию Set из 12го задания
	set := set.NewSet()

	//Проходим по символам всей строки и добавляем их во множество
	for _, v := range lower {
		set.Add(string(v))
	}

	setLen := set.Size()
	inLen := len(in)

	//Сравниваем длину исходной строки и множества, если одинаковые, то все элементы - уникальны
	return setLen == inLen
}

func main() {
	fmt.Printf("Unique of Aabcgd: %v\n", checkUnique("Aabcgd"))
	fmt.Printf("Unique of asdfghj: %v\n", checkUnique("asdfghj"))
}
