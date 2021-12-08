/*/* Имеется последовательность строк -
(cat, cat, dog, cat, tree) создать для нее собственное множество. */

//Реализация в set.go

package main

import (
	"fmt"

	"github.com/Ustelemov/WBLVL1/12/set"
)

func main() {
	in := []string{"cat", "cat", "dog", "tree", "cat"}
	set := set.NewSet()
	set.AddMul(in...)
	fmt.Println(set.ToSlice())
}
