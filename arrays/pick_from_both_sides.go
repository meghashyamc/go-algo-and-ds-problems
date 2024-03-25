package arrays

import (
	"math"
)

/*
Problem Description



Given an integer array A of size N.

You have to pick exactly B elements from either left or right end of the array A to get the maximum sum.

Find and return this maximum possible sum.

NOTE: Suppose B = 4 and array A contains 10 elements then

You can pick the first four elements or can pick the last four elements or can pick 1 from the front and 3 from the back etc. you need to return the maximum possible sum of elements you can pick.


Problem Constraints
1 <= N <= 10^5

1 <= B <= N

-10^3 <= A[i] <= 10^3



Input Format
First argument is an integer array A.

Second argument is an integer B.



Output Format
Return an integer denoting the maximum possible sum of elements you picked.



Example Input
Input 1:

 A = [5, -2, 3 , 1, 2]
 B = 3
Input 2:

 A = [1, 2]
 B = 1


Example Output
Output 1:

 8
Output 2:

 2


Example Explanation
Explanation 1:

 Pick element 5 from front and element (1, 2) from back so we get 5 + 1 + 2 = 8
Explanation 2:

 Pick element 2 from end as this is the maximum we can get

*/

func getMaxSumFromBothSides(arr []int, numOfElementsToPick int) int {

	if numOfElementsToPick == 1 {
		if arr[0] > arr[len(arr)-1] {
			return arr[0]
		}
		return arr[len(arr)-1]
	}
	sumFromLeftArray := []int{}
	sumFromRightArray := []int{}

	sum := 0
	for _, el := range arr {
		sum += el
		sumFromLeftArray = append(sumFromLeftArray, sum)
	}
	sum = 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum += arr[i]
		sumFromRightArray = append(sumFromRightArray, sum)
	}

	maxSumFromBothSides := math.MinInt

	for i := 0; i <= numOfElementsToPick-2; i++ {
		currentSum := sumFromLeftArray[i] + sumFromRightArray[numOfElementsToPick-2-i]

		if currentSum > maxSumFromBothSides {
			maxSumFromBothSides = currentSum
		}
	}

	return maxSumFromBothSides

}
