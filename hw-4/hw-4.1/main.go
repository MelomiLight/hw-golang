package main

import (
    "fmt"
   
)

var counter int
var slice []int


func appendToSlice(value int) {
    slice = append(slice, value)
}




func incrementCounter() {
    counter++
}
func decrementCounter() {
    counter--
}
/*Здесь две горутины одновременно пытаются манипулировать над значением, 
в итоге с разным количеством итерации разные результаты либо 0 либо 1 или даже 2
*/
func main() {
    for i:=0; i<20; i++ {
        go incrementCounter()
        go decrementCounter()
    }
    fmt.Println(counter)
// тут горутины стараются аппендить в слайс но в итоге из-за race condition результат всегда разный и явно не 1000
    for i := 0; i < 1000; i++ {
        go func(i int) {
            appendToSlice(i)
        }(i)
    }

    fmt.Println(len(slice)) // Результат может быть меньше 1000
}
