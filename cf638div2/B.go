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

func solve(n, k int) {

    arr := make([]int, n+5)
    truth_arr := make([]bool, 105)
    possible := true

    for i := 0; i < n; i += 1 {

        if i < n {
            Scanf("%d ", &arr[i])
        } else {
            Scanf("%d\n", &arr[i])
        }
        //set truth balues
        if i < k {
            truth_arr[arr[i]] = true
        } else if truth_arr[arr[i]] == false {
            possible = false
        }
    }

    Println(n, k, arr, possible)
}

func main() {
    defer Writer.Flush()

    var t int = 1
    Scanf("%d\n", &t)

    for ; t > 0; t -= 1 {
        var n, k int
        Scanf("%d %d\n", &n, &k)
        solve(n, k)

        // if !possible {
        //     Printf("-1\n")
        //     continue
        // }
        // // Println(arr, n, k)
        // //repeat the array in intervals of k
        //
        // ans := make([]int, 0, 10005)
        // //first copy first k elements
        // last_indx := 0
        // pointer := k
        // for i := 0; i < 10005; i += 1 {
        //
        //     if i < k {
        //         ans = append(ans, arr[i])
        //     } else {
        //         ans = append(ans, ans[i-k])
        //         if ans[i] == arr[pointer] {
        //             pointer += 1
        //         }
        //
        //         if pointer == n {
        //             last_indx = i
        //             break
        //         }
        //     }
        //     // Println(ans, pointer)
        // }
        // Printf("%d\n", len(ans))
        // for i := 0; i <= last_indx; i += 1 {
        //     if i < last_indx {
        //         Printf("%d ", ans[i])
        //     } else {
        //         Printf("%d\n", ans[i])
        //     }
        // }
    }
}
