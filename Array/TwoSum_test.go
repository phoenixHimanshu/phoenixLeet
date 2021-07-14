package main

import "testing"

//Input: nums = [2,7,11,15], target = 9
func TestTwoSum(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	targetSum := 9

	result := twoSum(nums, targetSum)
	println(result[0])
	println(result[1])
}

//Input:num[3,2,4], target = 6
func TestTwoSum1(t *testing.T) {
	nums := []int{3, 2, 4}
	targetSum := 6
	result := twoSum(nums, targetSum)
	println(result[0])
	println(result[1])
}

/*
Input: nums = [3,3], target = 6
Output: [0,1]
*/
func TestTwoSum2(t *testing.T) {
	nums := []int{3, 3}
	targetSum := 6
	result := twoSum(nums, targetSum)
	println(result[0])
	println(result[1])
}

// Input: nums = [0,4,3,0], target = 0
// Output: [0,3]

func TestTwoSum3(t *testing.T) {
	nums := []int{0, 4, 3, 0}
	targetSum := 0
	result := twoSum(nums, targetSum)
	println(result[0])
	println(result[1])
}

//[-1,-2,-3,-4,-5] -8
func TestTwoSum4(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	targetSum := -8
	result := twoSum(nums, targetSum)
	println(result[0])
	println(result[1])
}
