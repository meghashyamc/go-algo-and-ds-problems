package binarysearch

/*Given a sorted array, use binary search to search if an integer exists in the array.
If it exists, return an array with indices that hold that integer, else return an empty array.

Input Format
A sorted array and the integer being searched for.

Output Format
An array consisting of the indices of the integer being searched for or an empty array.

Example Input
Input 1:

	A = [2, 5, 9, 10, 12, 100]
	B = 9

Example Output
Output 1:

	[2]

*/

func search(arr []int, num int) []int {

	if arr == nil || len(arr) == 0 {
		return []int{}
	}

	return searchHelper(arr, 0, len(arr), num)
}

func searchHelper(arr []int, start, end, num int) []int {

	result := []int{}
	for {
		if start >= end {
			return result
		}
		mid := start + (end-start)/2

		if num < arr[mid] {
			end = mid
			continue
		}

		if num > arr[mid] {
			start = mid + 1
			continue
		}

		if num == arr[mid] {
			result = append(result, mid)
			for i := mid - 1; i >= 0; i-- {
				if num == arr[i] {
					result = append(result, i)
					continue
				}
				break
			}

			for i := mid + 1; i < len(arr); i++ {
				if num == arr[i] {
					result = append(result, i)
					continue
				}
				break
			}
			return result
		}

	}

}
