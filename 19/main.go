/*Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»).
Символы могут быть unicode. */
package main

import "fmt"

func reverseString(in string) string {
	runes := []rune(in) //получаем слайс рун, кодирующих один символ в unicode
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	out := string(runes) //переводим слайс рун в строку
	return out
}

func main() {
	fmt.Printf("Input: %s\n", "😎❤️🎃🔥")
	fmt.Printf("Reversed: %s\n", reverseString("😎❤️🎃🔥"))
}
