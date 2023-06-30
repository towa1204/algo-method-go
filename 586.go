package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	length := 1000 + 1
	S := make([]int, length)
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		S[a]++
	}

	for i := 0; i < M; i++ {
		var b int
		fmt.Scan(&b)
		S[b]++
	}

	for i := 1; i < length; i++ {
		if S[i] == 2 {
			fmt.Println(i)
		}
	}

}
