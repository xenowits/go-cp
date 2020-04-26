package main

import (
    "github.com/xenowits/go-cp/io"
    "math/rand"
    "time"
)

func main() {
    defer io.Writer.Flush()
    var t int = 1
    io.Scanln(&t)
    for t > 0 {
        arr := make([]float64, 0, 10)
        rand.Seed(time.Now().Unix())

        for i := 0; i < 5; i += 1 {
            arr = append(arr, 100*rand.NormFloat64())
        }

        io.Println(arr)

        t -= 1
    }
}
