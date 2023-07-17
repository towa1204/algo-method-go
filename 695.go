package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		A[i] = a
	}

	sort.Ints(A)

	var ans float64
	if N%2 == 0 {
		// 偶数のとき
		ans = float64((A[N/2-1] + A[N/2])) / float64(2)
	} else {
		// 奇数のとき
		ans = float64(A[(N-1)/2])
	}
	fmt.Println(ans)
}
