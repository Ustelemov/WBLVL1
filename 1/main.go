/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования). */

package main

import "fmt"

//так как набор остальных полей может быть произвольный, опустим их и создадим для Human один метод Talk,
//выводящий имя (значение поля name) из структуры Human

type Human struct {
	name string
}

func (h Human) Talk() {
	fmt.Printf("%s is talking\n", h.name)
}

//реализуем has-a подход, создав анонимное поле, встраювающее Human в структуру Action
type Action struct {
	Human
}

//теперь создав переменную типа Action и заполнив в Human поле name мы можем вызвать
//метод структуры Human от Action
func main() {
	act := Action{
		Human: Human{
			name: "Ivan",
		},
	}
	act.Talk()
}
