package math

import (
	"log"
	"strconv"
	"strings"
)

/* Problem Description

You are given two numbers A and B.

You have to add them without using arithmetic operators and return their sum.



Problem Constraints
1 <= A, B <= 10^9


Input Format
The first argument is the integer A. The second argument is the integer B.


Output Format
Return a single integer denoting their sum.


Example Input
Input 1:
A = 3
B = 10
Input 2:

A = 6
B = 1


Example Output
Output 1:
13
Output 2:

7


Example Explanation
Explanation 1:
3 + 10 = 13
Explanation 2:

6 + 1 = 7.
Note, you have to add without using arithmetic operators. */
const (
	valZero = "0"
	valOne  = "1"
)

type sumHelper struct {
	sum   string
	carry string
}

var binDigitSummations = map[string]sumHelper{"000": {sum: "0", carry: "0"}, "001": {sum: "1", carry: "0"}, "010": {sum: "1", carry: "0"}, "011": {sum: "0", carry: "1"}, "100": {sum: "1", carry: "0"}, "101": {sum: "0", carry: "1"}, "110": {sum: "0", carry: "1"}, "111": {sum: "1", carry: "1"}}

func addNumbersWithoutOperator(num1, num2 int) int {

	binStringForNum1 := strconv.FormatInt(int64(num1), 2)
	binStringForNum2 := strconv.FormatInt(int64(num2), 2)

	sumBinStrings := addBinStrings(binStringForNum1, binStringForNum2)

	result, err := strconv.ParseInt(sumBinStrings, 2, 64)
	if err != nil {
		log.Fatalf("Failed to parse the sum of the two numbers: %v", err)
	}

	return int(result)

}

func addBinStrings(s1, s2 string) string {

	lengthDiff := len(s2) - len(s1)
	if lengthDiff > 0 {
		s1 = prependZeroes(s1, lengthDiff)
	} else if lengthDiff < 0 {
		s2 = prependZeroes(s2, -lengthDiff)
	}

	var sum string
	currentSumDetails := sumHelper{sum: valZero, carry: valZero}
	for i := len(s1) - 1; i >= 0; i-- {
		currentSumDetails = binDigitSummations[string(s1[i])+string(s2[i])+currentSumDetails.carry]
		sum = currentSumDetails.sum + sum
		if i == 0 {
			sum = currentSumDetails.carry + sum
		}
	}

	return sum

}

func prependZeroes(s string, zeroesToPrepend int) string {

	var b strings.Builder
	for i := 0; i < zeroesToPrepend; i++ {
		b.WriteString(valZero)

	}
	b.WriteString(s)

	return b.String()
}
