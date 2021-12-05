/*Разработать программу,
которая переворачивает словав строке.Пример: «snow dog sun — sun dog snow». */

package main

import (
	"fmt"
	"strings"
)

func reverseWords(in string) string {
	//получаем слайс слов, разделенных во входной строке одним или несколькими пробельными символами
	fields := strings.Fields(in)
	n := len(fields)

	for i := 0; i < n/2; i++ {
		fields[n-1-i], fields[i] = fields[i], fields[n-1-i]
	}

	out := strings.Join(fields, " ") //переводим слайс слов в строку, разделяя слова пробелом

	return out
}

func main() {
	fmt.Printf("Input: %s\n", "❤️ ➡️ 🚀")
	fmt.Printf("Reversed: %s\n", reverseWords("❤️ ➡️ 🚀"))
}
