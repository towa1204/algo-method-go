package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	sum := 0
	for i := 1; i <= N-1; i++ {
		for j := i + 1; j <= N; j++ {
			sum += i * j
		}
	}
	fmt.Println(sum)
}
