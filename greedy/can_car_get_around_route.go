package greedy

/*
Given two integer arrays A and B of size N. There are N gas stations along a circular route, where the amount of gas at station i is A[i].

You have a car with an unlimited gas tank and it costs B[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the minimum starting gas station's index if you can travel around the circuit once, otherwise return -1.

You can only travel in one direction. i to i+1, i+2, ... n-1, 0, 1, 2.. Completing the circuit means starting at i and ending up at i again.


Problem Constraints
1 <= |A| <= 5 * 10^5
|A| == |B|
0 <= Ai <= 5 * 10^3
0 <= Bi <= 5 * 10^3


Input Format
The first argument given is the integer array A. The second argument given is the integer array B.


Output Format
Return the minimum starting gas station's index if you can travel around the circuit once, otherwise return -1.


Example Input
A = [1, 2]
B = [2, 1]


Example Output
1


Example Explanation
If you start from index 0, you can fill in A[0] = 1 amount of gas.
Now your tank has 1 unit of gas. But you need B[0] = 2 gas to travel to station 1.

If you start from index 1, you can fill in A[1] = 2 amount of gas.
Now your tank has 2 units of gas. You need B[1] = 1 gas to get to station 0.
So, you travel to station 0 and still have 1 unit of gas left over.
You fill in A[0] = 1 unit of additional gas, making your current gas = 2. It costs you B[0] = 2 to get to station 1, which you do and complete the circuit.


*/

func canCompleteCircuit(gasAvailableAtIndices []int, gasNeededFromThisIndexToNext []int) int {

	// If some of gas available at indices is less than the sum of gas needed from this index to next, then
	// we can't complete circuit from any index
	if sumUpArray(gasAvailableAtIndices)-sumUpArray(gasNeededFromThisIndexToNext) < 0 {
		return -1
	}

	netGasAvailable := 0
	for i, j := 0, 0; ; {

		netGasAvailable += gasAvailableAtIndices[j] - gasNeededFromThisIndexToNext[j]
		if netGasAvailable < 0 {
			netGasAvailable -= gasAvailableAtIndices[i] - gasNeededFromThisIndexToNext[i]

			j = getNextCircularIndex(len(gasAvailableAtIndices), j)
			i = j
			netGasAvailable = 0
			continue
		}
		if getNextCircularIndex(len(gasAvailableAtIndices), j) == i {

			return i
		}

		j = getNextCircularIndex(len(gasAvailableAtIndices), j)

	}

}

func getNextCircularIndex(totalIndices, currentIndex int) int {

	return (currentIndex + 1) % totalIndices

}

func sumUpArray(arr []int) int {

	sum := 0

	for _, num := range arr {

		sum += num
	}
	return sum
}
