package codility

func TapeEquilibrium(A []int) int {
	N := len(A)

	leftSums := make([]int, N)
	leftSums[0] = A[0]
	for i := 1; i < N-1; i++ {
		leftSums[i] = leftSums[i-1] + A[i]
	}

	rightSums := make([]int, N)
	rightSums[N-1] = A[N-1]
	minDiff := absi(leftSums[N-2] - rightSums[N-1])
	for i := N - 2; i > 0; i-- {
		rightSums[i] = rightSums[i+1] + A[i]
		if diff := absi(leftSums[i-1] - rightSums[i]); diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func absi(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

//fmt.Println("min:", Solution([]int{3, 1, 2, 4, 3}))
