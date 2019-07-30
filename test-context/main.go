package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func workder(ctx context.Context, wg *sync.WaitGroup, i int) error {
    defer wg.Done()

    for {
        select {
        default:
            fmt.Println(i, "goroutine监控中...")
            time.Sleep(1 * time.Second)
        case <-ctx.Done():
            fmt.Println(i, "监控退出，停止了...")
            return ctx.Err()
        }
    }
}

func main() {
    ctx := context.Background()
    ctx1, cancel := context.WithTimeout(ctx, 1 * time.Second)
    ctx2 := context.WithValue(ctx1, "kkk", "vvvvvvvv")
    fmt.Println("ctx: ", ctx)
    fmt.Println("ctx1: ", ctx1)
    fmt.Println("ctx2: ", ctx2)

    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go workder(ctx2, &wg, i)
    }

    time.Sleep(3 * time.Second)
    fmt.Println("可以了，通知监控停止")
    cancel()

    wg.Wait()
}