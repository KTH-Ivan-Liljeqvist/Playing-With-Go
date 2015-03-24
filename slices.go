/*
	@author - Ivan Liljeqvist
	@version - 23-03-2015
*/

package main

import "golang.org/x/tour/pic"

/*
	This program includes a function that creates a picture from a matrix using slices in Go.
*/

/*
	This function returns a matrix of dimensions passed in as parameters.
	The matrix is populated with elements derived from the formula: (x * 3 * y * 5) / 2.
*/

func Pic(dx, dy int) [][]uint8 {
	//return slice of length dy
	//each element - slice of dx 8-bit unsign ints

	//initialize the slice which will hold the image
	matrix := make([][]uint8, dx)

	//loop through and populate the matrix
	for x := 0; x < dx; x++ {
		//initialize the subarray
		matrix[x] = make([]uint8, dy)
		//populate the subarray
		for y := 0; y < dy; y++ {
			//use the formula to create a value for the matrix element
			matrix[x][y] = uint8((x * 3 * y * 5) / 2)
		}
	}

	return matrix

}

func main() {
	pic.Show(Pic)
}
