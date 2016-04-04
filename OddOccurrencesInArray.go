package main

import "fmt"

func Solution0(A []int) int {
	// write your code in Go 1.4
	for i := 0; i < len(A)-1; i += 2 {
		for j := i + 1; j < len(A); j++ {
			if A[j] == A[i] { // Swap with A[i+1]
				A[i+1], A[j] = A[j], A[i+1]
				break
			}
		}
	}
	return A[len(A)-1]
}

func Solution(A []int) int {
	var matchedCount = map[int]int{}
	for _, e := range A {
		matchedCount[e]++
	}
	for _, e := range A {
		if matchedCount[e]%2 == 1 {
			return e
		}
	}
	return -1
}

func main() {
	a := []int{1, 3, 5, 9}
	a[0], a[1] = a[1], a[0]
	fmt.Println(Solution(a))
}
