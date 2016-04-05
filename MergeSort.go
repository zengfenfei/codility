package main

import "fmt"

// Recursive. Top down
func mergeSort(A []int) []int {
	n := len(A)
	result := make([]int, n)
	if n == 2 { // swap the 2 elements
		if A[0] > A[1] {
			result[0], result[1] = A[1], A[0]
		} else {
			result[0], result[1] = A[0], A[1]
		}
	} else if n > 2 {
		part1 := A[:n/2]
		part2 := A[len(part1):]
		sorted1 := mergeSort(part1)
		sorted2 := mergeSort(part2)

		// Merging to result
		i, i1, i2 := 0, 0, 0
		for ; i1 < len(sorted1) && i2 < len(sorted2); i++ {
			if sorted1[i1] < sorted2[i2] {
				result[i] = sorted1[i1]
				i1++
			} else {
				result[i] = sorted2[i2]
				i2++
			}
		}
		restPart, restI := sorted1, i1
		if i2 < len(sorted2) {
			restPart, restI = sorted2, i2
		}
		for ; i < n; i++ {
			result[i] = restPart[restI]
			restI++
		}
	} else {
		result[0] = A[0]
	}
	return result
}

func main() {
	A := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println(mergeSort(A))
	fmt.Println(mergeSort([]int{3}))
	fmt.Println(mergeSort([]int{3, 2}))
	fmt.Println(mergeSort([]int{3, 2, 9}))
}
