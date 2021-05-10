package main

//febonacci  1 1 2 3 5 8 13 21

func febo(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return febo(n-1) + febo(n-2)
}

func main() {
	for i := 0; i < 10; i++ {
		println(febo(i))
	}
}
