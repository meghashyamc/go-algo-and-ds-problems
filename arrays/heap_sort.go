package arrays

/*Given an array, sort it using the heap sort algorithm.

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

func heapSort(arr []int) {

	buildMaxHeap(arr)

	heapLength := len(arr)
	for i := len(arr) - 1; i > 0; i-- {

		exchange(arr, 0, i)
		heapLength--
		maxHeapify(arr, 0, heapLength)

	}

}

func buildMaxHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, len(arr))
	}

}

func maxHeapify(arr []int, i, heapLength int) {

	if heapLength < 2 {
		return
	}
	if i > heapLength-1 {
		return
	}

	leftChildIndex := getLeftChildIndex(i)
	rightChildIndex := getRightChildIndex(i)

	indexWithLargestVal := getIndexWithLargestVal(arr, i, leftChildIndex, rightChildIndex, heapLength)

	if indexWithLargestVal == i {
		return
	}

	exchange(arr, i, indexWithLargestVal)
	maxHeapify(arr, indexWithLargestVal, heapLength)
	return

}

func getLeftChildIndex(i int) int {

	return 2*i + 1
}

func getRightChildIndex(i int) int {

	return 2*i + 2
}

func getIndexWithLargestVal(arr []int, i, j, k, heapLength int) int {

	indexWithLargestVal := i

	if j < heapLength && arr[j] > arr[indexWithLargestVal] {
		indexWithLargestVal = j
	}

	if k < heapLength && arr[k] > arr[indexWithLargestVal] {
		indexWithLargestVal = k
	}

	return indexWithLargestVal

}

func exchange(arr []int, i, j int) {

	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
