package main

import "fmt"

//К чему требуется адаптировать
type Animal interface {
	Speak()
}

type People struct{}

func (p People) Speak() {
	fmt.Println("Hello!")
}

//Адаптируемый
type Cat interface {
	Meow()
}

type BlackCat struct{}

func (c BlackCat) Meow() {
	fmt.Println("Meow!")
}

//Адаптер
type AdapterCat struct {
	cat Cat
}

func (a AdapterCat) Speak() {
	a.cat.Meow()
}

func main() {
	people := People{}
	adapterCat := AdapterCat{
		cat: BlackCat{},
	}

	listenAnimal(people)
	listenAnimal(adapterCat)
}

func listenAnimal(animal Animal) {
	animal.Speak()
}
