/* Имеется последовательность строк -
(cat, cat, dog, cat, tree) создать для нее собственное множество. */
package main

import "fmt"

//Реализуем тип Set с помощью структуры с полем list типа map, в котором ключ (string) - это элемент в Set,
//а значение (bool) - идентификатор наличия элемента
//Также реализуем несколько стандартных методов: вставка одного элемента, вставка нескольких элементов,
//проверка наличия элемента, перевод в слайс
type Set struct {
	list map[string]bool
}

func NewSet() *Set {
	return &Set{
		list: make(map[string]bool, 0),
	}
}

func (set *Set) Add(val string) {
	set.list[val] = true
}

func (set *Set) AddMul(vals ...string) {
	for _, val := range vals {
		set.list[val] = true
	}
}

func (set *Set) Has(val string) bool {
	return set.list[val]
}

//Перевод множества (Set) в Slice
func (set *Set) ToSlice() []string {
	out := make([]string, len(set.list))
	index := 0

	for k := range set.list {
		out[index] = k
		index += 1
	}
	return out
}

func main() {
	in := []string{"cat", "cat", "dog", "tree", "cat"}
	set := NewSet()
	set.AddMul(in...)
	fmt.Println(set.ToSlice())
}
