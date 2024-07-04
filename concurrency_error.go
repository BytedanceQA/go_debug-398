package main

import (
    "fmt"
    "sync"
)

var (
    counter int
    wg      sync.WaitGroup
)

func concurrency_error() {
    wg.Add(2)
    go increment("A")
    go increment("B")
    wg.Wait()
    fmt.Println("最终值:", counter)
}

func increment(name string) {
    for i := 0; i < 1000; i++ {
        counter++
    }
    fmt.Println("完成", name)
    wg.Done()
}