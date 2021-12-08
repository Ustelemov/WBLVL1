/*Удалить i-ый элемент из слайса. */
package main

import "fmt"

//Вариант без модификации старого (исходного) слайса модификацией нового (см. main)
//В варианте return append(in[:index], in[index+1:]) - такая модификация возможна
func RemoveIndex(s []int, index int) []int {
	out := make([]int, 0)
	out = append(out, s[:index]...)
	return append(out, s[index+1:]...)
}

func main() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("input: ", input) //[0 1 2 3 4 5 6 7 8 9]
	removeIndex := RemoveIndex(input, 5)

	fmt.Println("--Remove index 5--")
	fmt.Println("input: ", input)             //[0 1 2 3 4 5 6 7 8 9]
	fmt.Println("removeIndex: ", removeIndex) //[0 1 2 3 4 6 7 8 9]

	removeIndex[0] = 999
	fmt.Println("--Change index 0 to 999 in removeIndex result--")
	fmt.Println("input: ", input)             //[0 1 2 3 4 5 6 7 9 9]
	fmt.Println("removeIndex: ", removeIndex) //[999 1 2 3 4 6 7 8 9]
}
