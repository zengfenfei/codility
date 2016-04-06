package codility

func Distinct(A []int) int {
	MergeSort(A)
	count := 0
	for i, e := range A {
		if i == 0 || e != A[i-1] {
			count++
		}
	}
	return count
}
