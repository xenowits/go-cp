package main

import (
    "github.com/xenowits/go-cp/io"
)

func maxScore(s string) int {
    n := len(s)

    ones := make([]int, n)
    zeroes := make([]int, n)

    if s[0] == '1' {
        ones[0] = 1
    } else {
        zeroes[0] = 1
    }

    for i := 1; i < n; i += 1 {

        ones[i] = ones[i-1]
        zeroes[i] = zeroes[i-1]

        if s[i] == '1' {
            ones[i] += 1
        } else {
            zeroes[i] += 1
        }

        //io.Println(i, ones[i], zeroes[i])

    }

    var ans int = -1

    for i := 0; i < n-1; i += 1 {
        var temp int
        temp = zeroes[i] + (ones[n-1] - ones[i])
        if temp > ans {
            ans = temp;
        }
        //io.Println(i, temp, ans)
    }

    return ans
}

func main() {
    defer io.Writer.Flush()
    var s string
    io.Scanln(&s)
    x := maxScore(s)
    io.Println(x)
}
