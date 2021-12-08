/*Реализовать собственную функцию sleep. */
package main

import (
	"fmt"
	"time"
)

func Sleep(sec int) {
	//блокируемся до того, как пройдет sec секунд
	<-time.After(time.Second * time.Duration(sec))
}

func main() {
	fmt.Println("Before sleep...")
	Sleep(10)
	fmt.Println("After sleep...")
}
