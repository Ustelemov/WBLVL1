/*Реализовать все возможные способы остановки выполнения горутины. */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//1 вариант: приостановка горутины через некоторое время по чтению из канала таймера
	timer := time.NewTimer(3 * time.Second)
	go func() {
		fmt.Println("Gorute 1 started")
		for {
			select {
			case <-timer.C:
				fmt.Println("Gorute 1 is stopping")
				return
			default:
				_ = 0
				//Какие-то действия
			}
		}
	}()

	//2 Вариант: ожидание чтение из kill-канала, если значение появилось в канале,
	//нужно остановить выполнени

	kill := make(chan interface{})

	go func() {
		fmt.Println("Gorute 2 started")
		for {
			select {
			case <-kill:
				fmt.Println("Gorute 2 is stopping")
				return
			default:
				_ = 0
				//Какие-то действия
			}
		}
	}()

	//3 Вариант: также работаем с каналом kill из варианта 2, однако ожидаем не сообщение,
	//а проверяем закрыт ли канал и тогда завершаем горутину.
	//Более "идеоматичная" (правильная, характерная для го) реализация в варианте 6

	go func() {
		fmt.Println("Gorute 3 started")
		for {
			select {
			case _, ok := <-kill: //ok - true, если канал не закрыт, false - закрыт
				if !ok {
					fmt.Println("Gorute 3 is stopping")
					return
				}
			default:
				_ = 0
				//Какие-то действия
			}
		}
	}()

	//4 Вариант: c помощью контекста типа WithCancel, возвращающего функцию отмены
	ctxCancel, cancelCancel := context.WithCancel(context.Background())
	defer cancelCancel()

	go func() {
		fmt.Println("Gorute 4 started")
		for {
			select {
			case <-ctxCancel.Done(): //если cancel() была вызвана
				fmt.Println("Gorute 4 is stopping")
				return
			default:
				_ = 0
				//Какие-то действия
			}
		}
	}()

	//5 Вариант: c помощью контекста типа WithTimeout, по истечению заданного
	//времени сообщает, с помощью канала Done, о необходимости завершения работы
	ctxTimeOut, cancelTimeOut := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelTimeOut()

	go func() {
		fmt.Println("Gorute 5 started")
		for {
			select {
			case <-ctxTimeOut.Done(): //если timeout "истек"
				fmt.Println("Gorute 5 is stopping")
				return
			default:
				_ = 0
				//Какие-то действия
			}
		}
	}()

	//6 Вариант с помощью закрытия канала, из которого производиться чтение в range
	data := make(chan int)
	go func() {
		fmt.Println("Gorute 6 started")
		for d := range data {
			_ = d
		}
		fmt.Println("Gorute 6 is stopping")
	}()

	//7 Вариант: c помощью контекста типа WithDeadline, при превышении заданного времени
	//сообщает, с помощью канала Done, о необходимости завершения работы
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancelDeadline()

	go func() {
		fmt.Println("Gorute 7 started")
		for {
			select {
			case <-ctxDeadline.Done(): //если время превышено
				fmt.Println("Gorute 7 is stopping")
				return
			default:
				_ = 0
				//Какие-то действия
			}
		}
	}()

	//Немного подождем, сымитируем процесс
	time.Sleep(3 * time.Second)
	//Отправляем в kill канал сообщение, завершающее горутину варианта 2
	kill <- true
	//Закрываем канал, чем завершаем горутину варианта 3
	close(kill)
	//Вызываем функцию cancel, завершающую горутину варианта 4
	cancelCancel()
	//Закрываем канал из которого читались данные, чем завершаем горутину варианта 6
	close(data)

	//Дождемся вывода всех сообщений
	time.Sleep(5 * time.Second)
}
