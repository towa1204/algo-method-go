package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	ans := 0
	for i := 1; i <= N; i++ {
		ans += i
	}
	fmt.Println(ans)
}
