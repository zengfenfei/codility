package main

import "fmt"

func Solution(S string, P []int, Q []int) []int {
	M, N := len(P), len(S)
	//A, C, G and T
	aCounts := make([]int, N+1)
	cCounts := make([]int, N+1)
	gCounts := make([]int, N+1)

	for i := 0; i < N; i++ {
		aCounts[i+1] = aCounts[i]
		cCounts[i+1] = cCounts[i]
		gCounts[i+1] = gCounts[i]
		var counts []int
		switch S[i] {
		case 'A':
			counts = aCounts
		case 'C':
			counts = cCounts
		case 'G':
			counts = gCounts
		}
		if counts == nil {
			continue
		}
		counts[i+1]++
	}

	minFactors := make([]int, M)
	for i := 0; i < M; i++ {
		from, to := P[i], Q[i]+1
		if aCounts[to]-aCounts[from] > 0 {
			minFactors[i] = 1
		} else if cCounts[to]-cCounts[from] > 0 {
			minFactors[i] = 2
		} else if gCounts[to]-gCounts[from] > 0 {
			minFactors[i] = 3
		} else {
			minFactors[i] = 4
		}
	}
	return minFactors
}

func main() {
	fmt.Println("minFactors:", Solution("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}))
}
