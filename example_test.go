package codility_test

import (
	"fmt"
	. "github.com/zengfenfei/codility"
	"testing"
)

func ExampleFish() {
	A, B := make([]int, 5), make([]int, 5)
	A[0] = 4
	A[1] = 3
	A[2] = 2
	A[3] = 1
	A[4] = 5

	B[0] = 0
	B[1] = 1
	B[2] = 0
	B[3] = 0
	B[4] = 0
	fmt.Println(Fish(A, B))
	// Output: 2
}

func ExampleNumberOfDiscIntersections() {
	A := make([]int, 6)
	A[0] = 1
	A[1] = 5
	A[2] = 2
	A[3] = 1
	A[4] = 4
	A[5] = 0
	fmt.Println(NumberOfDiscIntersections(A))
	// Output: 11
}
func ExampleTriangle() {
	A := []int{10, 2, 5, 1, 8, 20}
	fmt.Println(Triangle(A))
	// Output: 1
}

func ExampleDistinct() {
	var A [6]int
	A[0] = 2
	A[1] = 1
	A[2] = 1
	A[3] = 2
	A[4] = 3
	A[5] = 1
	fmt.Println(Distinct(A[:]))
	// Output: 3
}

// ExampleMinAvgTwoSlice test and demonstrate MinAvgTwoSlice
func ExampleMinAvgTwoSlice() {
	var A [7]int
	A[0] = 4
	A[1] = 2
	A[2] = 2
	A[3] = 5
	A[4] = 1
	A[5] = 5
	A[6] = 8
	fmt.Println(MinAvgTwoSlice(A[:]))
	fmt.Println(MinAvgTwoSlice([]int{10000, -10000}))
	// Output:
	// 1
	// 0
}

func getSortCases() [][]int {
	return [][]int{
		{6, 5, 3, 1, 8, 7, 2, 4},
		{3},
		{3, 2},
		{3, 2, 9},
	}
}

func ExampleInsertSort() {
	for _, a := range getSortCases() {
		InsertSort(a)
		fmt.Println(a)
	}
	// Output:
	// [1 2 3 4 5 6 7 8]
	// [3]
	// [2 3]
	// [2 3 9]
}

func ExampleSelectionSort() {
	for _, a := range getSortCases() {
		SelectionSort(a)
		fmt.Println(a)
	}
	// Output:
	// [1 2 3 4 5 6 7 8]
	// [3]
	// [2 3]
	// [2 3 9]
}

func ExampleMergeSort() {
	for _, a := range getSortCases() {
		MergeSort(a)
		fmt.Println(a)
	}
	// Output:
	// [1 2 3 4 5 6 7 8]
	// [3]
	// [2 3]
	// [2 3 9]
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort([]int{6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4, 6, 5, 3, 1, 8, 7, 2, 4})
	}
}
