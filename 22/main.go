/*Разработать программу, которая перемножает, делит,складывает,
вычитает две числовых переменных a,b,
значение которых > 2^20 */

package main

import (
	"fmt"
	"math/big"
)

func main() {

	x1 := 1<<63 - 1
	x2 := 1<<63 - 1

	bigX1 := big.NewInt(int64(x1))
	bigX2 := big.NewInt(int64(x2))

	fmt.Printf("Sum: %d+%d = %d\n", x1, x2, new(big.Int).Add(bigX1, bigX2))
	fmt.Printf("Sub: %d-%d = %d\n", x1, x2, new(big.Int).Sub(bigX1, bigX2))
	fmt.Printf("Mul: %d*%d = %d\n", x1, x2, new(big.Int).Mul(bigX1, bigX2))
	fmt.Printf("Div: %d/%d = %d\n", x1, x2, new(big.Int).Div(bigX1, bigX2))

}
