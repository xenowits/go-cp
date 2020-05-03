package main

import (
    "fmt"
    "bufio"
    "os"
)

//FOR I/O : STARTS
var Reader *bufio.Reader = bufio.NewReader(os.Stdin)
var Writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func Printf (f string, a ...interface{}) { fmt.Fprintf(Writer, f, a...) }
func Scanf (f string, a ...interface{}) { fmt.Fscanf(Reader, f, a...) }
func Println (a ...interface{}) { fmt.Fprintln(Writer, a...)}
func Scanln (a ...interface{}) { fmt.Fscanln(Reader, a...)}
//I/O ENDS HERE

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func Min(x, y int64) int64 {
    if x < y {
        return x
    } else {
        return y
    }
}

func Max(x, y int64) int64 {
    if x >= y {
        return x
    } else {
        return y
    }
}

func Ceil(x, y int64) int64 {
    //returns x/y
    if x%y == 0 {
        return x/y
    } else {
        return int64(x/y)+1
    }
}

func main() {
    defer Writer.Flush()

    var t int = 1
    Scanf("%d\n", &t)

    for ; t > 0; t -= 1 {
        var n int
        Scanf("%d\n", &n)
        arr := make([]int, n)
        no_line := make([]bool, 11)
        ans := make([]int, 11)

        maxu := int(-1e7)
        minu := int(1e7)

        for i := 0; i < n; i += 1 {
            Scanf("%d ", &arr[i])
            no_line[arr[i]] = true
        }

        for i := 0; i < 11; i += 1 {
            ans_temp := 1
            if no_line[i] {
                for j := i
            }
        }

        iter := []int {-2, -1, 1, 2}

        // Println(iter, arr, m)

        for i := 0; i < n; i += 1 {



            if temp1 > maxu {
                maxu = temp1
            } else if temp1 < minu {
                minu = temp1
            }
        }

        Printf("%d %d\n", minu+1, maxu+1)

    }
}
