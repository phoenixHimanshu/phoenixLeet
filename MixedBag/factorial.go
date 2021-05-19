package main

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return fact(n-1) * n

}
