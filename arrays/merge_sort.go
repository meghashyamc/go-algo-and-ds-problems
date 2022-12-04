package arrays

import (
	"fmt"
	"strings"
)

/*Given an array, sort it using the merge sort algorithm.

Input Format
The first and only argument is an array.

Output Format
No output needs to be returned. The passed array should be sorted.

Example Input
Input 1:

	A = [15, 8, 9, 3, 5, 2]

Example Output
Output 1:

	[2, 3, 5, 9, 8, 15]

*/

func mergeSort(arr []int) {

	if arr == nil || len(arr) < 2 {
		return
	}
	mergeSortHelper(arr, 0, len(arr))
}

func mergeSortHelper(arr []int, start, end int) {

	if end-start <= 1 {
		return
	}

	mid := start + (end-start)/2
	mergeSortHelper(arr, start, mid)
	mergeSortHelper(arr, mid, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	fmt.Println(fmt.Sprintf("merge called for [%s] and [%s]", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr[start:mid])), ","), "[]"), strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr[mid:end])), ","), "[]")))

	leftArr := make([]int, mid-start)
	rightArr := make([]int, end-mid)
	if len(leftArr) == 0 || len(rightArr) == 0 {
		return
	}

	copy(leftArr, arr[start:mid])
	copy(rightArr, arr[mid:end])

	i, j := 0, 0

	for k := start; k < end; k++ {

		if i >= len(leftArr) {
			arr[k] = rightArr[j]
			j++
			continue
		}
		if j >= len(rightArr) {
			arr[k] = leftArr[i]
			i++
			continue
		}
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
			continue
		}
		arr[k] = rightArr[j]
		j++

	}

}
