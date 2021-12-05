/*Реализовать пересечение двух неупорядоченных множеств. */
package main

import (
	"fmt"
	"sort"
)

//В конце каждой функции к результату будем вызывать функцию, удаляющую дубликаты O(N).

//Вариант с проходом по одному множеству и для каждого элемента проход по всему второму множеству
//с проверкой есть ли проверяемый элемент из первого множества во втором
//O(N^2)
func simpleIntersect(set1 []int, set2 []int) []int {
	intersect := make([]int, 0)

	for _, v1 := range set1 {
		for _, v2 := range set2 {
			if v1 == v2 {
				intersect = append(intersect, v1)
				break
			}
		}
	}

	intersect = removeDuplicates(intersect)

	return intersect
}

//Использование map: добавляем все элементы перовго множества в map, а затем проверяем
//для каждого элемента второго множества наличие такого элемента в map
//O(N)
func hashMapIntersect(set1 []int, set2 []int) []int {
	intersect := make([]int, 0)
	m := make(map[int]bool)

	for _, v1 := range set1 {
		m[v1] = true
	}

	for _, v2 := range set2 {
		if m[v2] {
			intersect = append(intersect, v2)
		}
	}

	intersect = removeDuplicates(intersect)

	return intersect
}

//Отсортируем первое множество O(NlogN) и с помощью бинарного поиска O(logN) будем проверять
//Есть ли каждый элемент второго множества в отсортированном первом множестве
func sortIntersect(set1 []int, set2 []int) []int {
	intersect := make([]int, 0)
	sort.Ints(set1)

	for _, v := range set2 {
		i := sort.Search(len(set1), func(i int) bool { return set1[i] >= v })
		if i < len(set1) && set1[i] == v {
			intersect = append(intersect, v)
		}
	}

	intersect = removeDuplicates(intersect)

	return intersect
}

//Удаляем дубликаты (если это потребуется), функции intersect возвращают результат с дубликатами
//Используем подход с map, добавляем новые элементы в результирующее множество, в map, и устанавливаем им значение true
//Проверяем для каждого элемента входного множества наличие элемента в map, если есть - не добавляем (уже добавляли)
func removeDuplicates(in []int) []int {
	out := make([]int, 0)
	m := make(map[int]bool)

	for _, v := range in {
		if !m[v] {
			m[v] = true
			out = append(out, v)
		}
	}
	return out
}

func main() {
	//в качестве множеств будем использовать целые числа
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 9, 10}
	set2 := []int{2, 4, 8, 10, 50, 5, 8, 9, 10}

	fmt.Println(simpleIntersect(set1, set2))
	fmt.Println(hashMapIntersect(set1, set2))
	fmt.Println(sortIntersect(set1, set2))

}
