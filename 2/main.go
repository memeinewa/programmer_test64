package main

func Commission(n float32) (com float32) {
	if n <= 10000 {
		com = 0
	} else if n <= 20000 {
		com = n * 0.05
	} else if n > 20000 {
		com = n * 0.1
	}
	return com
}

func CompareCommission(a float32, b float32) float32 {
	return a - b
}
