package cachelookup

func rowwiseLookup() {
	var arr [10000][10000]int

	for t := 0; t < 10; t++ {
		for i := 0; i < 10000; i++ {
			for j := 0; j < 10000; j++ {
				arr[i][j] += 1
			}
		}
	}
}

func columnwiseLookup() {
	var arr [10000][10000]int

	for t := 0; t < 10; t++ {
		for i := 0; i < 10000; i++ {
			for j := 0; j < 10000; j++ {
				arr[j][i] += 1
			}
		}
	}
}
