package arrays

// Given a matrix of m * n elements (m rows, n columns), return all elements of the matrix in spiral order.

// Example:

// Given the following matrix:

// [
//     [ 1, 2, 3 ],
//     [ 4, 5, 6 ],
//     [ 7, 8, 9 ]
// ]
// You should return

// [1, 2, 3, 6, 9, 8, 7, 4, 5]

const (
	toright  = 0
	tobottom = 1
	toleft   = 2
	totop    = 3
)

func spiralOrder(A [][]int) []int {

	spiralOrderArr := make([]int, 0)
	rows := len(A)
	if rows == 0 {
		return spiralOrderArr
	}
	cols := len(A[0])
	if cols == 0 {
		return spiralOrderArr
	}

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	dir := toright
	var nexti, nextj int
	for i, j := 0, 0; ; {
		if i > rows-1 || j > cols-1 || i < 0 || j < 0 || visited[i][j] {
			dir, nexti, nextj = turnright(dir)
			i = i + nexti
			j = j + nextj

			if i > rows-1 || j > cols-1 || i < 0 || j < 0 || visited[i][j] {
				break
			}
			continue
		}

		spiralOrderArr = append(spiralOrderArr, A[i][j])
		visited[i][j] = true
		switch dir {
		case toright:
			j++
		case tobottom:
			i++
		case toleft:
			j--
		case totop:
			i--
		}
	}
	return spiralOrderArr
}

func turnright(dir int) (newdir int, nexti int, nextj int) {

	switch dir {
	case toright:
		return tobottom, 1, -1
	case tobottom:
		return toleft, -1, -1
	case toleft:
		return totop, -1, 1
	case totop:
		return toright, 1, 1
	}
	return -1, -1, -1
}
