package main

//TODO implimentation fails if s1 and s2 have different size
/*
	l(string1, string2) =   if c1 == c2      c1 and c2 are first char from string1 and string2
							1 + l( r1 + r2)
							if c1 != c2
							max( l(r1,s1),l(s1,r2))
*/

func longestCommonSubsequenceLentgh(s1 string, s2 string) int {

	/*
	 * s2 will be on the rows, s1 will be on the columns.
	 *
	 * +1 to leave room at the left for the "".
	 */

	cache := make([][]int, len(s1)+1)

	for indx := range cache {

		cache[indx] = make([]int, len(s2)+1)
	}

	/*
	 * cache[s2.length()][s1.length()] is our original subproblem.

	 *	Each entry in the
	 * table is taking a substring operation against whatever string is on the rows
	 *
	 *
	 * It goes from index 0 to index s2Row/s1Col (exclusive)
	 *
	 * So if my s1 = "azb" and s1Col = 2...then my substring that I pass to the
	 * lcs() function will be:
	 *
	 * 0 1 2 "a  z  b"
	 *
	 * "az" (index 2...our upper bound of the snippet...is excluded)
	 */

	for s2Row := 0; s2Row <= len(s2); s2Row++ {
		for s1Col := 0; s1Col <= len(s1); s1Col++ {
			if s2Row == 0 || s1Col == 0 {
				cache[s2Row][s1Col] = 0
			} else if string(s2[s2Row-1]) == string(s1[s1Col-1]) {
				cache[s2Row][s1Col] = cache[s2Row-1][s1Col-1] + 1
			} else {
				if cache[s2Row-1][s1Col] > cache[s2Row][s1Col-1] {
					cache[s2Row][s1Col] = cache[s2Row-1][s1Col]
				} else {
					cache[s2Row-1][s1Col] = cache[s2Row][s1Col-1]
				}

			}
		}

	}
	return cache[len(s2)][len(s1)]
}

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
