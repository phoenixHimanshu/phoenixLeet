package main

import "testing"

/*

Example 1:  Works

Input:
s1 = "ABCD"
s2 = "ABCD"

Output: 4
Explanation:
"ABCD"
"ABCD"
Both strings share the subsequence "A", "B", "C", "D".


Example 2:

Input:
s1 = "ADC"
s2 = "ABCD"

Output: 2  ==> Does not work
Explanation:
"ADC"
"ABCD"
Both strings share the subsequence "A", "D".


Example 3:  ==> Does not work

Input:
s1 = "DBC"
s2 = "CBD"

Output: 1
Explanation:
"DBC"
"CBD"
      or
"DBC"
"CBD"
      or
"DBC"
"CBD"

*/

func TestLCD(t *testing.T) {
	println(longestCommonSubsequenceLentgh("ABCD", "ABCD"))
}
