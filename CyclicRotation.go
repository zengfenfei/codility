package main

import "fmt"

func Solution(A []int, K int) []int {
	N := len(A)
	if N == 0 {
		return A
	}
	K = K % N
	rotated := make([]int, len(A))
	// Move N-K elements to the right
	for i := K; i < N; i++ {
		rotated[i] = A[i-K]
	}
	// Rotate the shifted out elements to left
	for i := 0; i < K; i++ {
		rotated[i] = A[N-K+i]
	}
	return rotated
}

func main() {
	fmt.Println(Solution([]int{3, 8, 9, 7, 6}, 3)) // [9, 7, 6, 3, 8]
	fmt.Println(Solution([]int{}, 3))              // [9, 7, 6, 3, 8]
}
