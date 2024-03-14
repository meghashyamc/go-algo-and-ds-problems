package arrays

/*
You are given an array of 0s and 1s in random order. Segregate 0s on left side and 1s on right side of the array
[Basically you have to sort the array].
Traverse array only once.
Problem Constraints
1<=|A|<=1e6

Input Format
First argument is array of integers consisting of 0's and 1's only.

Output Format
Return a sorted array.

Example Input
Input 1:
a=[0 1 0]
Input 2:

A=[1 1 0 ]

Example Output
Ouput 1:
[0 0 1]
Ouput 2:

[0 1 1]
*/
func segregateZeroesAndOnes(arr []int) []int {

	var zeroCount, oneCount int
	for _, el := range arr {
		if el == 0 {
			zeroCount++
			continue
		}
		oneCount++
	}

	for i := 0; i < zeroCount; i++ {
		arr[i] = 0
	}

	for i := zeroCount; i < zeroCount+oneCount; i++ {
		arr[i] = 1
	}

	return arr

}
