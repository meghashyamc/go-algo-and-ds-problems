package onstrings

import (
	"strings"
)

/*
Compare two version numbers version1 and version2.
If version1 > version2 return 1,
If version1 < version2 return -1,
otherwise return 0.
You may assume that the version strings are non-empty and contain only digits and the . character.
The . character does not represent a decimal point and is used to separate number sequences. For instance, 2.5 is not "two and a half" or "half way to version three", it is the fifth second-level revision of the second first-level revision.

Note: Here is an example of version numbers ordering:
0.1 < 1.1 < 1.2 < 1.13 < 1.13.4


Problem Constraints
1 <= |A|, |B| <= 5000


Input Format
The first argument is a string A representing version1.
The first argument is a string B representing version2.


Output Format
Return an integer.


Example Input
A = "1.13"

B = "1.13.4"



Example Output
-1


Example Explanation
Version1 = "1.13"
Version2 = "1.13.4"
Version1 < version2, hence return -1


*/

func compareVersionNumbers(version1, version2 string) int {

	versionsArray1 := strings.Split(version1, ".")

	versionsArray2 := strings.Split(version2, ".")

	for i := 0; ; i++ {

		if i > len(versionsArray1)-1 && i > len(versionsArray2)-1 {
			return 0
		}

		currentVersion2PositionValue := removePrefixZeroes(versionsArray2, i)

		if i > len(versionsArray1)-1 && len(currentVersion2PositionValue) > 0 {
			return -1
		}
		currentVersion1PositionValue := removePrefixZeroes(versionsArray1, i)
		if i > len(versionsArray2)-1 && len(currentVersion1PositionValue) > 0 {
			return 1
		}

		currentVersionsPositionResult := compareCurrentVersionPositionValues(currentVersion1PositionValue, currentVersion2PositionValue)
		if currentVersionsPositionResult != 0 {
			return currentVersionsPositionResult
		}
	}
}

func removePrefixZeroes(arr []string, i int) string {

	if i > len(arr)-1 {
		return ""
	}
	firstNonZeroIndex := 0
	for i, ch := range arr[i] {
		if ch == '0' {
			continue
		}
		firstNonZeroIndex = i
		break

	}

	if arr[i][firstNonZeroIndex:] == "0" {
		return ""
	}

	return arr[i][firstNonZeroIndex:]
}

func compareCurrentVersionPositionValues(version1PositionValue, version2PositionValue string) int {

	if len(version1PositionValue) > len(version2PositionValue) {
		return 1
	}

	if len(version1PositionValue) < len(version2PositionValue) {

		return -1
	}

	if version1PositionValue > version2PositionValue {

		return 1
	}

	if version1PositionValue < version2PositionValue {

		return -1
	}

	return 0

}
