package codility

import (
	"sort"
)

func Triangle(A []int) int {
	sort.Ints(A)
	for i := 2; i < len(A); i++ {
		if A[i-2]+A[i-1] > A[i] {
			return 1
		}
	}
	return 0
}
