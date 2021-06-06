package main

import "testing"

func TestAddTwoDigits(t *testing.T) {
	println(addTwoDigits(2, 3))
}

func TestAddTwoDigits1(t *testing.T) {
	println(addTwoDigitsRecursive(2, 3))
}

func TestHammingWeight(t *testing.T) {
	println(hammingWeight(3))
}
