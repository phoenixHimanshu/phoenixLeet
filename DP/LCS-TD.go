package main

func longestCommonSubsequenceLengthRecur(s1 string, s2 string) int {

	/*
	 * Base Case
	 *
	 * lcs("", anything...) == 0 lcs(anything..., "") == 0 lcs("", "") == 0
	 *
	 * A subproblem where either string is empty will have a result of 0. There can
	 * be nothing in common with an empty string and anything else.
	 */
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	/*
	 * Just extracting what we certainly will need and being explicit
	 */
	s1WithoutFinalCharacter := s1[0 : len(s1)-1]
	s2WithoutFinalCharacter := s2[0 : len(s2)-1]

	if string(s1[len(s1)-1]) == string(s2[len(s2)-1]) {
		/*
		 * No competition necessary. A match. The answer to THIS subproblem is 1 PLUS
		 * the best answer to the subproblem without either of these characters.
		 */
		return 1 + longestCommonSubsequenceLengthRecur(s1WithoutFinalCharacter, s2WithoutFinalCharacter)

	} else {

		/*
		 * Character mismatch. Compete subproblems, whoever wins becomes the answer to
		 * this subproblem and is hence returned
		 */
		option1 := longestCommonSubsequenceLengthRecur(s1, s2WithoutFinalCharacter)
		option2 := longestCommonSubsequenceLengthRecur(s1WithoutFinalCharacter, s2)

		if option1 > option2 {
			return option1
		} else {
			return option2
		}

	}

	//return 0
}
