package main

// LIS  Longest Increasing Subsequence

//2D array for memoization

// 1 2 5 3 4
func lengthOfLIS(arr []int) int {
	var memoo [10][10]int
	return LISHelper(-1, 0, arr, memoo)
}

func LISHelper(prev int, curr int, nums []int, memoo [10][10]int) int {
	var op1 int
	if int(curr) == len(nums) {
		return 0
	}

	if prev != -1 && memoo[prev][curr] != 0 {
		return memoo[prev][curr]
	}
	if prev == -1 || nums[prev] < nums[curr] {
		op1 = 1 + LISHelper(curr, curr+1, nums, memoo)
	}
	op2 := LISHelper(prev, curr+1, nums, memoo)

	if prev != -1 {
		if op1 > op2 {
			memoo[prev][curr] = op1
		} else {
			memoo[prev][curr] = op2
		}
	}

	if op1 > op2 {
		return op1
	} else {
		return op2
	}
}
