package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	sum := A[0] % 10
	for i := 1; i < N; i++ {
		sum = (sum * A[i]) % 10
	}
	fmt.Println(sum)
}
