package main

import "testing"

func TestLIS(t *testing.T) {
	arr := make([]int, 5)
	arr = []int{1, 2, 5, 3, 4}
	println(lengthOfLIS(arr))

}
