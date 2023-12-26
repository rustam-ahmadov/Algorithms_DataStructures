package main

// Constraints:

// 	n == matrix[0].length count of x coor
// 	m == matrix.length count of y coor

// 	1 <= m, n <= 200
// 	-231 <= matrix[i][j] <= 231 - 1

func setZeroes(matrix [][]int) {
	zeros := make(map[int]int)

	n := len(matrix[0])
	m := len(matrix)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeros[i*n+j] = 0
			}
		}
	}

	for pos := range zeros {
		y, x := pos/n, pos%n

		for i := 0; i < m; i++ {
			matrix[i][x] = 0
			if i == y {
				for j := 0; j < n; j++ {
					matrix[y][j] = 0
				}
			}
		}
	}
}

func main() {
	matrix := [][]int{{0,0,0,5},{4,3,1,4},{0,1,1,4},{1,2,1,3},{0,0,1,1}}
	setZeroes(matrix)
}
