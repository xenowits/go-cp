package main

import(
  "fmt"
)

const (
  MAXN = int(1e6+5)
)

var adj [][]int = make([][]int , MAXN)
var visited []bool = make([]bool, MAXN)

func getDegrees(n int){
  for i := 1; i <= n; i+=1 {
    fmt.Println(adj[i])
  }
}

func dfs(src int, par int) {
  fmt.Println(src, par)
  visited[src] = true
  for _, v := range adj[src] {
    if visited[v] == false {
      dfs(v, src)
    }
  }
}

func bfs(src int, nodes int) {
  visited := make([]bool, nodes+5)
  visited[src] = true
  var u int
  var que []int
  que = append(que, src)
  for len(que) > 0 {
    u = que[0]
    // fmt.Println(u)
    que = que[1:]
    for _, v := range adj[u] {
      if visited[v] == false {
        que = append(que,v)
        visited[v] = true
      }
    }
  }
}

func main() {
  var n, e, u, v int
  fmt.Scanf("%d %d", &n, &e)

  for i := 1; i <= e; i += 1 {
    fmt.Scanf("%d %d", &u, &v)
    adj[u] = append(adj[u], v)
    adj[v] = append(adj[v], u)
  }
  
  ch := make(chan int)
  bfs(1, n, ch)
  //dfs(1, -1)
}
