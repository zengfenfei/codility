package main

func Solution(X int, A []int) int {
	hasLeaves := make([]bool, X)
	validLeafCount := 0
	for i, e := range A {
		if e <= X && !hasLeaves[e-1] {
			validLeafCount++
		}
		if validLeafCount >= X {
			return i
		}
		hasLeaves[e-1] = true
	}
	return -1
}

func main() {

}
