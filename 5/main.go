/*Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Будем использовать буферезированный канал на одно значение, чтобы не блокироваться
	//при чтении
	channel := make(chan int, 1)
	var secsToStop int

	//Читаем с консоли число секунд, после которого программа должна завершиться
	fmt.Print("Input number of seconds to stop after: ")
	if _, err := fmt.Scan(&secsToStop); err != nil {
		fmt.Println("Bad number of seconds to stop after, program is shutting down")
		return
	}

	fmt.Println("Program started")

	//Создаем таймер, который отправит в канал сообщение, когда время истечет
	t := time.NewTimer(time.Second * time.Duration(secsToStop))

	for {

		//Проверяем закончилось ли время, чтением из канала
		//Делаем это в отдельном селекте, так как даже если время истечет
		//То может выбраться случайный - не тот кейс, и значения продолжат печататься
		select {
		case <-t.C:
			fmt.Printf("%d seconds if gone. Program is stopping\n", secsToStop)
			return
		default: //пустой default позволяет не блокироваться
		}

		select {
		case out := <-channel:
			fmt.Printf("Got from channel: %d\n", out)
		default:
			channel <- rand.Intn(100)
			//немного засыпаем после отправки значения в канал, чтобы сделать вывод читаемым
			time.Sleep(time.Millisecond * 500)
		}
	}
}
