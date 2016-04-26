package codility

// https://codility.com/programmers/task/missing_integer/
func MissingInteger(A []int) int {
	positiveMin := 0
	for _, e := range A {
		if e > 0 && (positiveMin == 0 || e < positiveMin) {
			positiveMin = e
		}
	}
	if positiveMin == 0 || positiveMin > 1 {
		return 1
	}

	n := len(A)
	possibleMissings := make([]bool, n-1)
	for _, e := range A {
		minOffset := e - positiveMin - 1
		if minOffset >= 0 && minOffset < len(possibleMissings) {
			possibleMissings[minOffset] = true
		}
	}
	for i, e := range possibleMissings {
		if !e {
			return positiveMin + i + 1
		}
	}
	return positiveMin + n
}
