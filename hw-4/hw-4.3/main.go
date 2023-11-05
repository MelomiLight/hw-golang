package main

import (
	"fmt"
	"sync"
)
//Бесконечный цикл, который добавляет числа от 0 до 99 в map m
func writeLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			mux.Lock() //чтобы заблокировать доступ к map
			m[i] = i
			mux.Unlock() //чтобы разблокировать доступ
		}
	}
}
//Бесконечный цикл, который читает и выводит пары ключ-значение из map m
func readLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock() //чтобы разрешить параллельное чтение map
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mux.RUnlock()//чтобы снять блокировку чтения
	}
}


func main() {
	m := map[int]int{}

	mux := &sync.RWMutex{} //экземпляр мьютекса sync.RWMutex 

	go writeLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)

	/*Для предотвращения завершения программы использую канал block, который ожидает, 
	чтобы программа не завершилась до его закрытия (<-block)
	*/
	block := make(chan struct{})
	<-block
}

