/*Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}. */
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	types := make([]interface{}, 4)
	types[0] = 10
	types[1] = "test"
	types[2] = true
	types[3] = make(chan int)

	fmt.Println("--Using switch type--")
	for _, v := range types {
		switch t := v.(type) {
		case int:
			fmt.Printf("int: %v\n", t)
		case string:
			fmt.Printf("string: %v\n", t)
		case bool:
			fmt.Printf("bool: %v\n", t)
		case chan int:
			fmt.Printf("chan int: %v\n", t)
		}
	}

	fmt.Println("--Using reflect.TypeOf()--")
	for _, v := range types {
		fmt.Printf("%s: %v\n", reflect.TypeOf(v), v)
	}

	fmt.Println(strconv.Quote("--Using fmt %T--"))
	for _, v := range types {
		fmt.Printf("%T: %v\n", v, v)
	}
}
