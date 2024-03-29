package hashing

/*
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Example:

Given [100, 4, 200, 1, 3, 2],

The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.

Your algorithm should run in O(n) complexity.

*/

func getLengthOfLongestConsecutiveSequence(arr []int) int {

	if len(arr) == 0 {
		return 0
	}
	longestSequenceLength := 1
	consideredNums := map[int]bool{}
	numsInArrMap := map[int]int{}

	for _, num := range arr {
		numsInArrMap[num]++

	}
	for _, num := range arr {

		if consideredNums[num] {
			continue
		}
		consideredNums[num] = true
		currentSequenceLength := 1

		for i := 1; ; i++ {

			if numsInArrMap[num+i] > 0 {
				consideredNums[num+i] = true
				currentSequenceLength++
				if len(arr)-len(consideredNums) == 0 {
					break
				}
			} else {
				break
			}

		}

		for i := 1; ; i++ {

			if numsInArrMap[num-i] > 0 {
				consideredNums[num-i] = true
				currentSequenceLength++
				if len(arr)-len(consideredNums) == 0 {
					break
				}
			} else {
				break
			}

		}

		if currentSequenceLength > longestSequenceLength {
			longestSequenceLength = currentSequenceLength
		}

	}
	return longestSequenceLength
}
