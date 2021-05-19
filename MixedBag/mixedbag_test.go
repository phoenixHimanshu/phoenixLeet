package main

import "testing"

func TestMixed(t *testing.T) {
	//Test 1
	println(fact(5))
	println("==============================")
	//Test 2  Print Fibonacci series
	for i := 0; i < 10; i++ {
		println(febo(i))
	}
}
