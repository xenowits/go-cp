package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'maxScore' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER m
 */

func maxScore(a []int32, mo int32) int32 {
    n := len(a)
    arr := make([]int, n)

    for i := 0; i < n; i += 1 {
        arr[i] = int(a[i])
    }

    //sort the slice
    sort.Ints(arr)

    fmt.Println(mo)

    m := int(mo)
    var no_of_times int = len(a)/m

    var ans, seg_ans, mod int64 = 0, 0, 1e9+7

    for i := 0; i < no_of_times; i += 1 {

        seg_ans = 0

        //starts from i*m

        for j := i*m; j < i*m+m; j += 1 {
            seg_ans = (seg_ans + int64(arr[j]))%mod
        }

        ans = (ans + int64(i+1)*seg_ans)%mod

        //fmt.Println(ans, seg_ans)
    }

    seg_ans = 0

    for i := m*no_of_times; i < n; i += 1 {

        seg_ans = (seg_ans + int64(arr[i]))%mod

    }

    ans = (ans + int64(no_of_times)*seg_ans)%mod

    return int32(ans)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    mo := int32(mTemp)

    aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var a []int32

    for i := 0; i < int(n); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    ans := maxScore(a, mo)

    fmt.Fprintf(writer, "%d\n", ans)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
