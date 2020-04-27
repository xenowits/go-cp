package main

import (
    "fmt"
    "bufio"
    "os"
)

var Reader *bufio.Reader = bufio.NewReader(os.Stdin)
var Writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func Printf (f string, a ...interface{}) { fmt.Fprintf(Writer, f, a...) }
func Scanf (f string, a ...interface{}) { fmt.Fscanf(Reader, f, a...) }
func Println (a ...interface{}) { fmt.Fprintln(Writer, a...)}
func Scanln (a ...interface{}) { fmt.Fscanln(Reader, a...)}

func main() {
    defer Writer.Flush()
    var t int = 1
    Scanln(&t)
    var s string
    m := make(map[string]int)
    for t > 0 {
         Scanln(&s)
        m[s] += 1
        t -= 1
    }
    cnt := 0
    for _, _ = range m {
        cnt += 1
    }
    Println(cnt)
}
