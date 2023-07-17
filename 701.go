package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	dist := make([]int, 5)

	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		if a <= 20 {
			dist[0]++
		} else if a <= 40 {
			dist[1]++
		} else if a <= 60 {
			dist[2]++
		} else if a <= 80 {
			dist[3]++
		} else {
			dist[4]++
		}
	}

	for _, d := range dist {
		fmt.Println(d)
	}
}
