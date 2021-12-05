/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика. */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int32
}

//функция инкремента
func (c *Counter) inc() int32 {
	return atomic.AddInt32(&c.value, 1)
}

//функция получения значения счетчика
func (c *Counter) get() int32 {
	return atomic.LoadInt32(&c.value)
}

//"Конструктор". Если потребуется, параметром можно передавать начальное значение счетчика
func NewCounter() *Counter {
	return &Counter{}
}

func main() {
	wg := sync.WaitGroup{}
	counter := NewCounter()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.inc()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Result: %d\n", counter.get())
}
