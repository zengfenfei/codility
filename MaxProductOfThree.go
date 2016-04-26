package codility

import (
	"sort"
)

// https://codility.com/programmers/task/max_product_of_three/
func MaxProductOfThree(A []int) int {
	sort.Ints(A)
	n := len(A)
	maxProduct := A[n-1] * A[n-2] * A[n-3] // Product of the largest 3 elements
	p := A[0] * A[1] * A[n-1]              // Product of the largest element and the smallest 2 elements
	if p > maxProduct {
		maxProduct = p
	}
	return maxProduct
}
