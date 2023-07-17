package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	sum := 0
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
	}

	avg := float64(sum) / float64(N)
	fmt.Println(avg)
}
