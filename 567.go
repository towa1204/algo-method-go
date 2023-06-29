package main

import "fmt"

func main() {
	var L, R int
	fmt.Scan(&L, &R)

	ans := 0
	for i := L; i <= R; i++ {
		ans += (2*i - 1) * (2*i - 1)
	}
	fmt.Println(ans)
}
