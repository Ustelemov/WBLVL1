/* Реализовать конкурентную запись данных в map. */
package main

import (
	"sync"
)

//Реализуем конкурентный map с помощью ConcurentMap
//и реализуем методы чтения и записи
//Тестировать безопасность будем при помощи тестовой функции (в main_test) и флага -race

type ConcurrentMap struct {
	mutex sync.RWMutex
	m     map[string]int
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{m: make(map[string]int)}
}

func (cm *ConcurrentMap) Load(key string) (int, bool) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	val, ok := cm.m[key]
	return val, ok
}

//Если элемент уже существует, то будет заменен
func (cm *ConcurrentMap) Store(key string, value int) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.m[key] = value
}
