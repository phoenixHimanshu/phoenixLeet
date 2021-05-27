package main

import "testing"

//Recursive Top Down Approch
func TestLIS(t *testing.T) {
	arr := make([]int, 5)
	arr = []int{1, 2, 5, 3, 4}
	println(lengthOfLIS(arr))

}

// This is Bottom up approch using iterative method
func TestBottomUp(t *testing.T) {
	//arr := make([]int, 5)
	arr := [25]int{1, 2, 5, 3, 4}
	println(lisIterative(arr))

}
func TestBottomUp1(t *testing.T) {
	arr := [25]int{10, 22, 9, 33, 21, 50, 41, 60, 80, 3}
	println(lisIterative(arr))
}
func TestTemparvari(t *testing.T) {
	println(longestCommonSubsequenceLentgh("DBC", "CBD"))
}
