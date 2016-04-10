package codility

// https://codility.com/programmers/task/perm_check/
func PermCheck(A []int) int {
	// write your code in Go 1.4
	n := len(A)
	occurs := make([]bool, n)
	for _, e := range A {
		if e < 1 || e > n {
			return 0
		}
		if occurs[e-1] {
			return 0
		}
		occurs[e-1] = true
	}
	return 1
}
