package main

func Factorials(n int) (sum int) {
	sum = 1
	for i := 2; i <= n; i++ {
		sum *= int(i)
	}
	return sum
}
