package main

import (
    //"github.com/xenowits/go-cp/io"
    //"math"
)

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

type node struct {
    r int
    c int
    sum int
}

func findDiagonalOrder(nums [][]int) []int {
    n := len(nums)
    var ans []int

    max_sum := 200005

    arr := make([]node, 0, max_sum)

    for i := 0; i < n; i += 1 {
        for j := 0; j < len(nums[i]); j += 1 {
            arr = append(arr, node {i, j, i+j})
        }
    }

    //sort according to sum and then acc to r

    sort.Slice(arr, func(i, j int) bool {
        if arr[i].sum == arr[j].sum {
            return arr[i].r > arr[j].r
        }
        return arr[i].sum < arr[j].sum
    })

    for _, i := range arr {
        //fmt.Println(i, i.r)
        ans = append(ans, nums[i.r][i.c])
    }

    fmt.Println(arr)

    return ans
}

func main() {
    defer Writer.Flush()
    var t int = 1
    Scanln(&t)

    for t > 0 {
    }
}
