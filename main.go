package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	s := make([]int, n+1)
	s[0] = n

	var input int
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		s[i+1] = input
	}

	sum := 0
	for _, v := range s {
		sum += v
	}

	fmt.Println(sum)
}
