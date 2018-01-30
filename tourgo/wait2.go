package main

import (
        "fmt"
        "sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var sum = 0

func process(n string) {
        wg.Add(1)
        go func() {
                defer wg.Done()

                for i := 0; i < 10000; i++ {
                        mu.Lock()
                        sum = sum + 1
                        mu.Unlock()
                }

                fmt.Println("From " + n + ":", sum)
        }()
}

func main() {
        processes := []string{"A", "B", "C", "D", "E"}
        for _, p := range processes {
                process(p)
        }

        wg.Wait()
        fmt.Println("Final Sum:", sum)
}
//sync.waitgroup cung cap 3 method la add, done va wait. add de xac dinh co bao nhieu gotutine can duoc cho. 
// done de ket thuc viec doi. when a gorutine exit, we must call done.
//he main goroutine blocks on Wait, Once the counter becomes 0, the Wait will return, and main goroutine can continue to run.
