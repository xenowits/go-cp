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
func Print (a ...interface{}) { fmt.Fprint(Writer, a...)}
func Scanln (a ...interface{}) { fmt.Fscanln(Reader, a...)}
//I/O ENDS HERE

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func Min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}

func main() {
    defer Writer.Flush()

    var t int = 1
    Scanf("%d\n", &t)

    for tc := 1; tc <= t; tc += 1 {

        var x, y int
        var str string
        Scanf("%d %d %s\n", &x, &y, &str)

        time_ := 0
        ans := int(1e7)

        for i := 0; i < len(str); i += 1 {

            time_needed := (Abs(x)+Abs(y))
            if time_needed <= time_ {
                //i can meet him here
                ans = Min(ans, time_)
            }

            // Println(x, y, time_, "ko")

            //now update them
            switch ch := str[i]; ch {
                case 'N':
                    y += 1
                case 'S':
                    y -= 1
                case 'E':
                    x += 1
                case 'W':
                    x -= 1
            }

             time_ += 1
        }

        // fmt.Println(x, y, time_, " sdf")

        time_needed := (Abs(x)+Abs(y))
        if time_needed <= time_ {
            //i can meet him here
            ans = Min(ans, time_)
        }

        if ans == int(1e7) {
            Printf("Case #%d: %s\n", tc, "IMPOSSIBLE")
        } else {
            Printf("Case #%d: %d\n", tc, ans)
        }
        // Println(x, y, str)

    }
}
