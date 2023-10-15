package main

import (
	"fmt"
	"sync"
	"time"
)

var start time.Time

func init() {
    start = time.Now()
}
/*
Этот пример я взял со статьи, здесь мы создаем два канала дальше используя конструкцию select ожидаем значение для чтения из канала, 
но так как нет никаких активных горутин выходит deadlock 
*/ 

func myFunction() error {
    start := time.Now()

    fmt.Println("myFunction() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    select {
    case res := <-chan1:
        fmt.Println("Response from chan1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from chan2", res, time.Since(start))
    case <-time.After(5 * time.Second): // Ожидаем результаты не более 5 секунд
        return fmt.Errorf("Deadlock detected")
    }

    fmt.Println("myFunction() stopped", time.Since(start))

    
    return nil
}

/*В этом коде функция взаимоблокировки создает ситуацию, когда две подпрограммы пытаются найти один и тот же мьютекс. 
Затем используется время, спящий режим позволяет подождать короткое время, а затем проверить, 
не произошла ли взаимоблокировка, используя select с тайм-аутом. Если обнаружена взаимоблокировка, 
он возвращает сообщение об ошибке.
*/
func myFunction2() error {
    var mu sync.Mutex

    go func() {
        mu.Lock()
        defer mu.Unlock()
        fmt.Println("Goroutine 1: Locked")
    }()

    go func() {
        mu.Lock() 
        defer mu.Unlock()
        fmt.Println("Goroutine 2: Locked")
    }()

   
    time.Sleep(2 * time.Second)

    select {
    case <-time.After(1 * time.Second):
        return nil // 
    default:
        return fmt.Errorf("Deadlock detected")
    }
}



func main() {
    if err := myFunction(); err != nil {
        fmt.Println("myFunction() error:", err)
    } else {
        fmt.Println("myFunction() completed successfully")
    }

	err := myFunction2()
    if err != nil {
        fmt.Println("DeadlockFunction() error:", err)
    } else {
        fmt.Println("DeadlockFunction() completed successfully")
    }
}