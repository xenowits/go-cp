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

const (
    MAXN = 1000006
)

func test(adj [][]int, n int) {

    //Slices hold references to an underlying array,
    //and if you assign one slice to another, both refer to the same array.
    //If a function takes a slice argument, changes it makes to the elements of the
    //slice will be visible to the caller, analogous to passing a pointer to the underlying array.
    //A Read function can therefore accept a slice argument rather than a pointer and a count;

    //perform BFS on the graph

    visited := make([]bool, n+5)
    visited[2] = true   //start of the BFS here
    visited[1] = true   //effectively breaks the link passing through 1

    var u int
    var que []int
    que = append(que, 2)

    // io.Println("graph")
    // for i := 1; i <= n; i++ {
    //     io.Println(i, adj[i])
    // }

    for len(que) > 0 {
      u = que[0]
      // io.Println(u)
      que = que[1:]
      for _, v := range adj[u] {
        if visited[v] == false {
          que = append(que,v)
          visited[v] = true
        }
      }
    }

    //count no of non-visited vertices

    var ans []int

    for _, i := range adj[1] {
        // io.Println(i, visited[i])
        if visited[i] == false {
            ans = append(ans, i)
        }
    }

    if len(ans) == 0 {
        io.Println(-1)
    } else {
        for i := 0; i < len(ans); i += 1 {
            io.Printf("%d ", ans[i])
        }
        io.Printf("\n")
        // io.Println(ans)
    }

}

func main() {
    defer io.Writer.Flush()
    var t int = 1
    io.Scanln(&t)

    for t > 0 {

        var n, m int
        io.Scanf("%d %d\n", &n, &m)

        //reserve memory for the graph
        var adj [][]int = make([][]int , MAXN)

        //input the edges
        for i := 0; i < m; i += 1 {
            var a, b int
            io.Scanf("%d %d\n", &a, &b)
            adj[a] = append(adj[a], b)
            adj[b] = append(adj[b], a)
        }

        test(adj, n)

        t -= 1

    }
}
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

const (
    MAXN = 1000006
)

func test(adj [][]int, n int) {

    //Slices hold references to an underlying array,
    //and if you assign one slice to another, both refer to the same array.
    //If a function takes a slice argument, changes it makes to the elements of the
    //slice will be visible to the caller, analogous to passing a pointer to the underlying array.
    //A Read function can therefore accept a slice argument rather than a pointer and a count;

    //perform BFS on the graph

    visited := make([]bool, n+5)
    visited[2] = true   //start of the BFS here
    visited[1] = true   //effectively breaks the link passing through 1

    var u int
    var que []int
    que = append(que, 2)

    // io.Println("graph")
    // for i := 1; i <= n; i++ {
    //     io.Println(i, adj[i])
    // }

    for len(que) > 0 {
      u = que[0]
      // io.Println(u)
      que = que[1:]
      for _, v := range adj[u] {
        if visited[v] == false {
          que = append(que,v)
          visited[v] = true
        }
      }
    }

    //count no of non-visited vertices

    var ans []int

    for _, i := range adj[1] {
        // io.Println(i, visited[i])
        if visited[i] == false {
            ans = append(ans, i)
        }
    }

    if len(ans) == 0 {
        io.Println(-1)
    } else {
        for i := 0; i < len(ans); i += 1 {
            io.Printf("%d ", ans[i])
        }
        io.Printf("\n")
        // io.Println(ans)
    }

}

func main() {
    defer io.Writer.Flush()
    var t int = 1
    io.Scanln(&t)

    for t > 0 {

        var n, m int
        io.Scanf("%d %d\n", &n, &m)

        //reserve memory for the graph
        var adj [][]int = make([][]int , MAXN)

        //input the edges
        for i := 0; i < m; i += 1 {
            var a, b int
            io.Scanf("%d %d\n", &a, &b)
            adj[a] = append(adj[a], b)
            adj[b] = append(adj[b], a)
        }

        test(adj, n)

        t -= 1

    }
}
