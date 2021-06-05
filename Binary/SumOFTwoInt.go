package main

//leetcode 371
//Given two integers a and b, return the sum of the two
//integers without using the operators + and -

func addTwoDigits(x int, y int) int {
	// Exit of loop is reached when carry becomes 0
	for y != 0 {
		// & operator is to get the carry positions
		carry := x & y
		// XOR operation simulates addition
		x = x ^ y
		// carry is applied one position left to where it was found
		// just like normal addition
		y = carry << 1

	}
	return x
}

/*
	8 bit machine representation
	x=2     0	0	0	0	0	0	1	0
	y=3		0	0	0	0	0	0	1	1

carry:=x&y  0	0	0	0	0	0	1	0
x = x^y		0	0	0	0	0	0	0	1
y=carry<<1  0	0	0	0	0	1	0	0

carry=x&y	0	0	0	0	0	0	0	0
x=x^y		0	0	0	0	0	1	0	1	//Answer
y=carry<<1	0	0	0	0	0	0	0	0	// loop breaks

*/

func addTwoDigitsRecursive(x int, y int) int {
	if y == 0 {
		return x
	} else {
		return addTwoDigitsRecursive(x^y, x&y<<1)
	}

}
