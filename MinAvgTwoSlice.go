package main

import "fmt"

func Solution(A []int) int {
	N := len(A)
	// There must be some 2 or 3 slices have minimal average
	minAvg := float64(A[0]+A[1]) / 2
	minAvgStart := 0

	if N > 2 {
		first3Avg := float64(A[0]+A[1]+A[2]) / 3
		if first3Avg < minAvg {
			minAvg = first3Avg
		}
	}
	for i := 1; i < N-2; i++ {
		avg := float64(A[i]+A[i+1]) / 2
		if avg < minAvg {
			minAvg = avg
			minAvgStart = i
		}
		avg = float64(A[i]+A[i+1]+A[i+2]) / 3
		if avg < minAvg {
			minAvg = avg
			minAvgStart = i
		}
	}
	// Check the last 2 elements
	if N > 2 {
		last2Avg := float64(A[N-2]+A[N-1]) / 2
		if last2Avg < minAvg {
			minAvg = last2Avg
			minAvgStart = N - 2
		}
	}
	return minAvgStart
}

func main() {
	var A [7]int
	A[0] = 4
	A[1] = 2
	A[2] = 2
	A[3] = 5
	A[4] = 1
	A[5] = 5
	A[6] = 8
	fmt.Println(Solution(A[:]))

	fmt.Println(Solution([]int{10000, -10000}))

}
