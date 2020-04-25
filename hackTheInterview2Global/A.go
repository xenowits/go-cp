package main

import (
	"fmt"
	"bufio"
	"sort"
	"os"
)

const (
	mod int64 = 1e9+7
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {

	defer writer.Flush()

	var n, m int
	scanf("%d %d\n", &n, &m)
	var no_of_times int = n/m

	arr := make([]int, n)
	for i := 0; i < n; i += 1 {
		scanf("%d ", &arr[i])
	}

	//sort the slice
	sort.Ints(arr)

	var ans, seg_ans int64 = 0, 0

	for i := 0; i < no_of_times; i += 1 {

		seg_ans = 0

		//starts from i*m

		for j := i*m; j < i*m+m; j += 1 {
			seg_ans = (seg_ans + int64(arr[j]))%mod
		}

		ans = (ans + int64(i+1)*seg_ans)%mod

		//print(ans, seg_ans)
	}

	seg_ans = 0

	for i := m*no_of_times; i < n; i += 1 {

		seg_ans = (seg_ans + int64(arr[i]))%mod

	}

	ans = (ans + int64(no_of_times)*seg_ans)%mod

	printf("%d\n",ans)
}
