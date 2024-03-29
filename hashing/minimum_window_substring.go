package hashing

/*

Given a string A and a string B, find the window with minimum length in A, which will contain all the characters in B in linear time complexity.
 Note that when the count of a character c in B is x, then the count of c in the minimum window in A should be at least x.

Note:

If there is no such window in A that covers all characters in B, return the empty string.
If there are multiple such windows, return the first occurring minimum window ( with minimum start index and length )

Problem Constraints
1 <= size(A), size(B) <= 10^6



Input Format
The first argument is a string A.
The second argument is a string B.



Output Format
Return a string denoting the minimum window.



Example Input
Input 1:

 A = "ADOBECODEBANC"
 B = "ABC"
Input 2:

 A = "Aa91b"
 B = "ab"




Example Output
Output 1:

 "BANC"
Output 2:

 "a91b"


Example Explanation
Explanation 1:

 "BANC" is a substring of A which contains all characters of B.
Explanation 2:

 "a91b" is the substring of A which contains all characters of B.



*/

func getMinWindowSubstring(str1, str2 string) string {

	if len(str1) < len(str2) {
		return ""
	}

	mapOfStr2Chars := map[rune]int{}
	for _, ch := range str2 {
		mapOfStr2Chars[ch]++
	}

	// Impossible values of substring start and end are initially chosen so we can
	// detect if str1 does not contain all characters of str2 even once
	minWindowSubstringStart := -1
	minWindowSubstringEnd := len(str1)
	var jIncrementedInPreviousIteration bool

	mapOfStr2CharsInStr1 := map[rune]int{}

	for i, j := 0, 0; i < len(str1) && j < len(str1); {
		charForJ := rune(str1[j])
		if _, ok := mapOfStr2Chars[charForJ]; !ok {
			j++
			jIncrementedInPreviousIteration = true
			continue
		}

		if jIncrementedInPreviousIteration || j == 0 {
			mapOfStr2CharsInStr1[charForJ]++
		}

		if isSecondMapIncludedInFirst(mapOfStr2CharsInStr1, mapOfStr2Chars) {

			currWindowSubstringLength := j - i + 1

			if currWindowSubstringLength < minWindowSubstringEnd-minWindowSubstringStart+1 {
				minWindowSubstringStart = i
				minWindowSubstringEnd = j
			}

			charForI := rune(str1[i])

			if _, ok := mapOfStr2Chars[charForI]; ok {
				mapOfStr2CharsInStr1[charForI]--
			}
			i++
			jIncrementedInPreviousIteration = false

			continue
		}

		j++
		jIncrementedInPreviousIteration = true

	}

	if minWindowSubstringStart < 0 {
		return ""
	}
	return str1[minWindowSubstringStart : minWindowSubstringEnd+1]

}

func isSecondMapIncludedInFirst(map1, map2 map[rune]int) bool {

	if len(map1) != len(map2) {
		return false
	}

	for k, v := range map1 {

		if map2[k] > v {
			return false
		}
	}
	return true
}
