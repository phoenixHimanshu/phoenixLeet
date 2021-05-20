package main

//Sample input  10 22 9 33 21 50 41 60 80 3

// Sample input 1 2 5 3 4
// N Square complexity
func lisIterative(arr [25]int) int {
	//Overall Maximum subset
	omax := 0
	var dp [10]int
	//First element in the array is single
	// so longest subsequence for fist element would be 1
	dp[0] = 1
	//Moving ahead from 2nd element
	for i := 1; i < len(dp); i++ {
		max := 0
		//Iterate upto ith element to check ith element
		// is greater than jth element
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				if dp[j] > max {
					max = dp[j]
				}
			}
		}
		// add one to previous max to become longest subsequence
		dp[i] = max + 1
		//if dp[i] is bigger then we have new winner
		if dp[i] > omax {
			omax = dp[i]
		}
	}
	return omax
}
