/*Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0 */

package main

import "fmt"

//Если требуется установить bit мы используем 1 << position
//Чтобы задать число с установленной в требуемой позиции 1 и остальными 0
//Отсчет позиции с нуля
//Затем применяем OR, которая оставляет все биты числа без изменения, а нужное
//устанавливает в 1: 0|1 = 1; 1|0 = 1; 1|1 = 1; 0|0 = 0
func setBit(number int64, position uint) int64 {
	number |= (1 << position)
	return number
}

//Для очистки бита (установления 0) создаем число с 1 в нужной позиции и инвертируем его
//^1000 = 0111. Используем AND, которая оставляет все биты без изменений, а нужны переводит в 0.
//1&0 = 0; 0&1 = 0; 1&1 = 1; 0&0 = 0
func clearBit(number int64, position uint) int64 {
	number &= ^(1 << position)
	return number
}

func main() {
	var number int64
	//Читаем число
	fmt.Print("Input number: ")
	if _, err := fmt.Scan(&number); err != nil {
		fmt.Println("Bad number input, program is shutting down")
		return
	}

	var bit uint

	//Читаем бит, который требуется установить
	fmt.Print("Input bit to set (0 or 1): ")
	if _, err := fmt.Scan(&bit); err != nil {
		fmt.Println("Bad bit input, program is shutting down")
		return
	}

	if bit > 1 {
		fmt.Println("Bad bit input (more 1), program is shutting down")
		return
	}

	var position uint

	//Читаем позицию, в которую требуется установить бит
	fmt.Print("Input position of change (count from 0): ")
	if _, err := fmt.Scan(&position); err != nil {
		fmt.Println("Bad position input, program is shutting down")
		return
	}

	if position > 63 {
		fmt.Println("Bad position input: can't be more than 63 (in int64), program is shutting down")
		return
	}

	if bit == 0 {
		fmt.Printf("Changed value: %d\n", clearBit(number, position))
	} else {
		fmt.Printf("Changed value: %d\n", setBit(number, position))
	}
}
