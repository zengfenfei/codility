package main

import "fmt"

func Solution(A int, B int, K int) int {
	// write your code in Go 1.4
	var start int
	r := A % K
	if r == 0 {
		start = A
	} else {
		start = A + K - r
		if start > B {
			return 0
		}
	}
	return 1 + (B-start)/K
}

func main() {
	check(11, 19, 17) // 20
	check(11, 40, 17) // 20
}

func check(A, B, K int) int {
	got := Solution(A, B, K)
	n := 0
	for i := A; i <= B; i++ {
		if i%K == 0 {
			n++
		}
	}
	if got == n {
		fmt.Println("ok", A, B, K)
	} else {
		fmt.Println("fail: expect", n, "got", got, "for", A, B, K)
	}
	return n
}
