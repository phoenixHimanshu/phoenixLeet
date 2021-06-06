package main

//Leetcode 191
//flip the rightmost 1

func hammingWeight(num uint32) uint32 {

	var count uint32
	if num == 0 {
		count = 0
	} else {
		count = 1
	}
	num = num & (num - 1)
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

/*


 */
