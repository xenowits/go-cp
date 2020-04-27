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

// Min returns the smaller of x or y.
func Min(x, y int64) int64 {
        if x > y {
                return y
        }
        return x
}

// Max returns the larger of x or y.
func Max(x, y int64) int64 {
        if x < y {
                return y
        }
        return x
}

func Abs(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    defer Writer.Flush()
    var t int = 1
    Scanln(&t)
    for t > 0 {
        var x, y, a, b, ans, temp, Y int64
        Scanf("%d %d\n%d %d\n", &x, &y, &a, &b)
        //Println(x,y,a,b)

        temp = Abs(x-y)
        Y = Min(x,y)

        ans = a*temp
        ans += Min(2*a*Y, b*Y)

        Println(ans)

        t -= 1
    }
}
