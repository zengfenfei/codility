package codility

func OddOccurrencesInArray(A []int) int {
	odd := 0
	for _, e := range A {
		odd ^= e
	}
	return odd
}

/*func main() {
	a := []int{1, 3, 5, 9}
	a[0], a[1] = a[1], a[0]
	fmt.Println(Solution(a))
}
*/
