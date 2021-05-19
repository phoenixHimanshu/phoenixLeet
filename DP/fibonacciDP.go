package main

//Fibonacci  1 1 2 3 5 8 13 21

var memo = make([]int, 10)

func febo(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = febo(n-1) + febo(n-2)
	return memo[n]
}

func main() {
	for i := 0; i < 10; i++ {
		println(febo(i))
		//println(febo(6))
	}

}
