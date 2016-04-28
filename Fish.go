package codility

// https://codility.com/programmers/task/fish/
func Fish(A []int, B []int) int {
	n := len(A)
	count := n
	downStack := make([]int, 0)
	for i := 0; i < n; i++ {
		if B[i] == 1 {
			downStack = append(downStack, A[i]) // stack push
		} else {
			j := len(downStack) - 1
			// eats downstream fish, downstack pop
			for ; j >= 0 && downStack[j] < A[i]; j-- {
				count--
			}
			// if downstream fish is not eaten up, then the current upstream fish must be eaten
			if j >= 0 {
				count--
			}
			downStack = downStack[:j+1]
		}
	}
	return count
}
