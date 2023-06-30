package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	A := make([]int, N)
	B := make([]int, M)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&B[i])
	}

	sum := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			sum += (A[i] + B[j])
		}
	}

	fmt.Println(sum)

}
