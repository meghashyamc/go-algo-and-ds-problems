package hashing

/*
Given a string A consisting of lowercase characters.

Check if characters of the given string can be rearranged to form a palindrome.

Return 1 if it is possible to rearrange the characters of the string A such that it becomes a palindrome else return 0.



Problem Constraints
1 <= |A| <= 10^5

A consists only of lower-case characters.



Input Format
First argument is an string A.



Output Format
Return 1 if it is possible to rearrange the characters of the string A such that it becomes a palindrome else return 0.



Example Input
Input 1:

 A = "abcde"
Input 2:

 A = "abbaee"


Example Output
Output 1:

 0
Output 2:

 1


Example Explanation
Explanation 1:

 No possible rearrangement to make the string palindrome.
Explanation 2:

 Given string "abbaee" can be rearranged to "aebbea" to form a palindrome.
*/

func checkPalindrome(s string) int {

	letterFrequencies := map[rune]int{}

	for _, ch := range s {
		letterFrequencies[ch]++
	}

	numOfOddFrequencies := 0
	for _, freq := range letterFrequencies {

		if freq%2 != 0 {
			numOfOddFrequencies++
		}
	}

	if numOfOddFrequencies <= 1 {
		return 1
	}
	return 0
}
