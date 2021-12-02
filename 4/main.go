/* Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров. */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Реализуем воркера, принимающего канал, из которого он будет читать данные и печатать их
//Исходя из реализации range данные читаются из канала до того момента, пока канал не будет закрыт.
//Для реализации такого воркера (который требуется по задаче) завершение его работы
//c помощью закрытия канала, это самый удобный\короткий\читабельный вариант, с помощью
//которого можно добиться требуемого результата

func worker(workerNum int, data <-chan interface{}) {
	fmt.Printf("worker %d started\n", workerNum)
	for d := range data {
		fmt.Printf("worker %d: %v\n", workerNum, d)
	}
}

func main() {

	var numWorkers int

	//Читаем с консоли количество воркеров
	fmt.Print("Input number of workers: ")
	if _, err := fmt.Scan(&numWorkers); err != nil {
		fmt.Println("Bad number of workers value, program is shutting down")
		return
	}

	done := make(chan os.Signal, 1) //канал для системного сигнала
	data := make(chan interface{})  //канал типа interface{}, так как данные - произвольные

	//SIGINT - сигнал прерывания программы (когда нажимаем Ctrl+C)
	//SIGTERM - сигнал завершения процесса
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < numWorkers; i++ {
		go worker(i+1, data)
	}
	for {
		select {
		case <-done: //проверяем пришел ли сигнал прирывания
			close(data)
			close(done)
			fmt.Println("\nWork finished")
			return
		default: // если нет - печатаем число и засыпаем, чтобы сделать вывод не сильно быстрым
			data <- rand.Intn(100) // в качестве данных для записи в канал генерируем случайное целое число [0,100)
			time.Sleep(300 * time.Millisecond)
		}
	}
}
