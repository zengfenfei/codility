package codility

//https://en.wikipedia.org/wiki/Sorting_algorithm#

// Recursive. Top down
func MergeSort(A []int) {
	B := make([]int, len(A)) // Temp
	mergeSortRecursive(A, B)
}

func mergeSortRecursive(A, B []int) {
	n := len(A)
	if n > 1 {
		splitMiddle := n / 2
		part1 := A[:splitMiddle]
		part2 := A[splitMiddle:]
		mergeSortRecursive(part1, B[:splitMiddle])
		mergeSortRecursive(part2, B[splitMiddle:])

		// Merging to result
		i1, i2 := 0, 0
		for i := 0; i < n; i++ {
			if i1 < splitMiddle && (i2 >= n-splitMiddle || part1[i1] < part2[i2]) {
				B[i] = part1[i1]
				i1++
			} else {
				B[i] = part2[i2]
				i2++
			}
		}
		for i := 0; i < n; i++ {
			A[i] = B[i]
		}
	}
}

func InsertSort(A []int) {
	n := len(A)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0 && A[j] > A[j+1]; j-- {
			A[j], A[j+1] = A[j+1], A[j]
		}
	}
}

func SelectionSort(A []int) {
	for i := 0; i < len(A); i++ {
		minI := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[minI] {
				minI = j
			}
		}
		A[i], A[minI] = A[minI], A[i]
	}
}
