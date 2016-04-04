package main

import "fmt"

func Solution(N int, A []int) []int {
	seq := make([]int, N)
	max := 0
	lastMax := 0
	for _, e := range A {
		if e >= 1 && e <= N {
			se := seq[e-1]
			if se < lastMax {
				se = lastMax
			}
			se++
			if se > max {
				max = se
			}
			seq[e-1] = se
		} else if e == N+1 {
			lastMax = max
		} else {
			// Invalid operation
		}
	}

	for i, e := range seq {
		if e < lastMax {
			seq[i] = lastMax
		}
	}
	return seq
}

func main() {
	a := []int{1, 3}
	a[3]++
	fmt.Println("a0", a[0])
}
