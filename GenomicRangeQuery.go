package codility

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	M, N := len(P), len(S)
	types := [...]byte{'A', 'C', 'G', 'T'}
	factors := [...]int{1, 2, 3, 4}
	ntypes := len(types)
	indexMap := make(map[byte]int, ntypes) // letter to prefixCounts index
	for i := 0; i < ntypes; i++ {
		indexMap[types[i]] = i
	}

	//A, C, G and T
	prefixCounts := make([][]int, ntypes-1) //Prefix counts of A, C, G
	for i := 0; i < len(prefixCounts); i++ {
		prefixCounts[i] = make([]int, N+1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < len(prefixCounts); j++ {
			prefixCounts[j][i+1] = prefixCounts[j][i]
		}
		index := indexMap[S[i]]
		if index < len(prefixCounts) {
			prefixCounts[index][i+1]++
		}
	}

	minFactors := make([]int, M)
	for i := 0; i < M; i++ {
		from, to := P[i], Q[i]+1
		for j, f := range factors {
			minFactors[i] = f
			if j < len(prefixCounts) && prefixCounts[j][to]-prefixCounts[j][from] > 0 {
				break
			}
		}
	}
	return minFactors
}

//	fmt.Println("minFactors:", Solution("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}))
