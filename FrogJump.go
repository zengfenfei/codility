package main

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	distance := Y - X
	steps := distance / D
	if distance%D == 0 {
		return steps
	} else {
		return steps + 1
	}
}

func main() {
	fmt.Println("steps", Solution())
}
