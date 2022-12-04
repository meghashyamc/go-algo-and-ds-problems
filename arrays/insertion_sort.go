package arrays

/*Given an array, sort it using the insertion sort algorithm.

Input Format
The first and only argument is an array.

Output Format
No output needs to be returned. The passed array should be sorted in place.

Example Input
Input 1:

	A = [15, 8, 9, 3, 5, 2]

Example Output
Output 1:

	[2, 3, 5, 9, 8, 15]

*/

func insertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for j := 1; j < len(arr); j++ {

		key := arr[j]
		i := j - 1
		for ; i >= 0 && arr[i] > key; i-- {

			arr[i+1] = arr[i]

		}
		arr[i+1] = key

	}

}
