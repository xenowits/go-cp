package main

import (
    "github.com/xenowits/go-cp/io"
    // "math/rand"
    // "time"
)

func main() {
    defer io.Writer.Flush()
    var t int = 1
    io.Scanln(&t)
    for t > 0 {

        t -= 1
    }
}
