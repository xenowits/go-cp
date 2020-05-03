package main

import (
    "fmt"
    "bufio"
    "os"
    "sort"
)

//FOR I/O : STARTS
var Reader *bufio.Reader = bufio.NewReader(os.Stdin)
var Writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func Printf (f string, a ...interface{}) { fmt.Fprintf(Writer, f, a...) }
func Scanf (f string, a ...interface{}) { fmt.Fscanf(Reader, f, a...) }
func Println (a ...interface{}) { fmt.Fprintln(Writer, a...)}
func Scanln (a ...interface{}) { fmt.Fscanln(Reader, a...)}
//I/O ENDS HERE

func Max(x, y int) int {
    if x < y {
        return y
    } else {
        return x
    }
}

func main() {
    defer Writer.Flush()

    var t, tc int = 1, 0
    Scanln(&t)

    ans := make([]string, t)

    for tc < t {
        var n, k int
        Scanf("%d %d\n", &n, &k)

        arr := make([]int, n+1)
        sorted_arr := make([][]int, k)

        max_len := int(-1e10)

        for i := 0; i < n; i += 1 {
            Scanf("%d ", &arr[i])
            sorted_arr[i%k] = append(sorted_arr[i%k], arr[i])
            max_len = Max(max_len, len(sorted_arr[i%k]))
        }

        for i := 0; i < k; i += 1 {
            sort.Ints(sorted_arr[i])
        }

        prev, indx := -1, 0
        can_be_sorted := true

        for indx < max_len {
            for i := 0; i < k; i += 1 {
                if indx < len(sorted_arr[i]) {
                     if sorted_arr[i][indx] > prev {
                         prev = sorted_arr[i][indx]
                     } else {
                         can_be_sorted = false
                         break;
                     }
                }
            }
            if !can_be_sorted {
                break
            }
            indx += 1
        }

        //Println(prev, indx, can_be_sorted)

        if can_be_sorted {
            // Println("yes")
            ans[tc] = "yes"
        } else {
            // Println("no")
            ans[tc] = "no"
        }

        tc += 1
    }

    for _, x := range ans {
        Println(x)
    }

}
