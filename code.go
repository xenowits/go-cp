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

func Abs(x int) int { if x < 0 { return -x } else { return x } }
func Min(x, y int) int { if x < y { return x } else { return y } }
func Max(x, y int) int { if x >= y { return x } else { return y } }
func Ceil(x, y int) int { if x%y == 0 { return x/y } else { return int(x/y)+1 } }

func main() {
    defer Writer.Flush()

    var t int = 1
    // Scanf("%d\n", &t)

    for ;t > 0; t -=1 {

    }
}
