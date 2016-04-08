package codility

// https://codility.com/programmers/task/binary_gap/
// N [1..2,147,483,647]
func BinaryGap(N int) int {
	max0s := 0
	cur0s := -1
	for i := uint8(0); i < 31; i++ {
		if N&(1<<i) > 0 { // Bit is 1
			if cur0s > max0s {
				max0s = cur0s
			}
			cur0s = 0
		} else if cur0s != -1 {
			cur0s++
		}
	}
	return max0s
}
