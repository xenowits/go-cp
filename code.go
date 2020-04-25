package main

import (
    "fmt"
    "bufio"
    "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }
func println(a ...interface{}) { fmt.Fprintln(writer, a...)}
func scanln(a ...interface{}) { fmt.Fscanln(reader, a...)}

func main() {
    defer writer.Flush()
    var a int
    scanln(&a)
    println(a)
}
