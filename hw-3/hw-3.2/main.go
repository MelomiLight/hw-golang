package main

import (
    "fmt"
    "sync"
)
/*
mergeChannels создает новый канал merged, в котором будут объединены данные из входных каналов channels.
Для каждого входного канала создается горутина, которая читает данные из него и отправляет их в merged.
Горутина для закрытия merged запускается после того, как все горутины, читающие из входных каналов, завершатся.
В функции main создаются три входных канала ch1, ch2 и ch3, и горутины отправляют значения в эти каналы.
Затем вызывается mergeChannels, и объединенные значения выводятся на экран в цикле.
*/
func mergeChannels(channels ...<-chan int) <-chan int {
    merged := make(chan int)

    var wg sync.WaitGroup
    wg.Add(len(channels))

    for _, ch := range channels {
        go func(ch <-chan int) {
            defer wg.Done()
            for num := range ch {
                merged <- num
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(merged)
    }()

    return merged
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

    go func() {
        defer close(ch1)
        ch1 <- 1
        ch1 <- 2
    }()

    go func() {
        defer close(ch2)
        ch2 <- 3
        ch2 <- 4
    }()

    go func() {
        defer close(ch3)
        ch3 <- 5
        ch3 <- 6
    }()

    merged := mergeChannels(ch1, ch2, ch3)

    for num := range merged {
        fmt.Println(num)
    }
}
