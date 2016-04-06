package codility

func ParsingCars(A []int) int {
	// write your code in Go 1.4
	var num0s, numPairs int
	for _, e := range A {
		if e == 0 {
			num0s++
		} else {
			numPairs += num0s
			if numPairs > 1e9 {
				return -1
			}
		}
	}
	return numPairs
}
