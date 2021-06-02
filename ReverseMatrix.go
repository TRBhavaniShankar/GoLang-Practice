package main

import "fmt"

func main() {

	// given example
	mainHelper(3, 3)

	// some more examples
	mainHelper(3, 4)
	mainHelper(3, 1)
	mainHelper(1, 4)
	mainHelper(3, 2)

	// constrants
	mainHelper(3, 0)
	mainHelper(0, 0)
	mainHelper(0, 2)

}

// function to reverse the matrix
// m : number of rows
// n : number of columns
// Traversing matrix complexity : O(mXn)
func reverseMatrix(matrix [][]int, m int, n int) (reversedMatrix [][]int) {

	// traverse the matrix column wise
	for i := 0; i < n; i++ {
		var rowToCol []int

		// traverse from the last row backwords so that we get 90 degree rotation
		for j := m - 1; j >= 0; j-- {
			rowToCol = append(rowToCol, matrix[j][i])
		}

		// at this point of iteration we have the column converted to row
		// therefore append the row to the new matrix
		reversedMatrix = append(reversedMatrix, rowToCol)
	}
	return reversedMatrix
}

// function to genreate matrix
// m : number of rows
// n : number of columns
func generateMatrix(m int, n int) (matrix [][]int) {

	// values starting from 1 to 20
	var temp int = 1
	for i := 0; i < m; i++ {

		//make rows
		var row []int
		for j := 0; j < n; j++ {

			// append values to rows
			row = append(row, temp)

			// increment the value for the next row item
			temp++

			// if item is more than 20 set the value back to 1
			// this is for simplicity
			if temp > 20 {
				temp = 1
			}

		}

		// append the row to the matrix
		matrix = append(matrix, row)

	}

	return matrix

}

// main helper to print the original and reversed matrix
// m : number of rows
// n : number of columns
func mainHelper(m int, n int) {

	if m == 0 || n == 0 {
		fmt.Println("Number of rows or columns is Zero")
		fmt.Println("=======================")
		return
	}

	var matrix [][]int = generateMatrix(m, n)
	fmt.Println("Original Matrix:")
	for i := 0; i < m; i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("-----------------------")

	var reversedMatrix [][]int = reverseMatrix(matrix, m, n)
	fmt.Println("Reversed matrix:")
	for i := 0; i < n; i++ {
		fmt.Println(reversedMatrix[i])
	}
	fmt.Println("=======================")

}
