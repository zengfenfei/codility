package codility

import (
	"sort"
)

// sort by the begin points
type Disc struct {
	Begin, End int
}

type SortableDiscs []Disc

func (d SortableDiscs) Len() int {
	return len(d)
}

func (d SortableDiscs) Less(i, j int) bool {
	return d[i].Begin < d[j].Begin
}

func (d SortableDiscs) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func NewDiscs(A []int) []Disc {
	discs := make([]Disc, len(A))
	for i := range discs {
		discs[i].Begin = i - A[i]
		discs[i].End = i + A[i]
	}
	return discs
}

// https://codility.com/programmers/task/number_of_disc_intersections/
func NumberOfDiscIntersections(A []int) int {
	discs := SortableDiscs(NewDiscs(A))
	sort.Sort(discs)
	pairs := 0
	for i := 0; i < len(discs)-1; i++ {
		// Binary search for the end
		for left, right := i+1, len(discs)-1; ; {
			if discs[i].End >= discs[right].Begin {
				pairs += right - i
				if pairs > 10000000 {
					return -1
				}
				break
			}
			mid := (left + right) / 2
			if discs[i].End >= discs[mid].Begin && mid != left {
				left = mid
			} else if mid != right {
				right = mid
			} else {
				break
			}
		}
	}
	return pairs
}
