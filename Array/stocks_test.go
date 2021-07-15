package main

import "testing"

func TestStock(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	println(result)
}

//[7,6,4,3,1]

func TestStock1(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	result := maxProfit(prices)
	println(result)
}
