/*
	@author - Ivan Liljeqvist
	@version - 23-03-2015
*/

package main

import (
	"fmt"
	"math"
)

/*
	This program includes a function that calculates the square root of the value passed in.
*/

/*
	This function returns the square root of the argument passed in.
*/

func Sqrt(x float64) float64 {

	//when the Z-value changes less than this constant, the loop will break
	const BREAK_THRESHOLD = 0.000001

	//the maximum number of iterations the loop will do if it's not broken
	const MAX_ITERATIONS = 100000

	//initial Z value
	var z = 1.0
	var lastZ = z

	//iteration counter
	var itCount = 0

	//loop through and recalculate Z
	for ; itCount < MAX_ITERATIONS; itCount++ {

		//recalculate Z
		z = z - (math.Pow(z, 2)-x)/(2*z)

		//see if it changed less than the threshold constant
		if (math.Abs(lastZ - z)) < BREAK_THRESHOLD {
			break
		}

		lastZ = z
	}

	//print the number of iterations

	fmt.Printf("Number of iterations: %v", itCount)
	fmt.Println("")
	fmt.Println("---")
	fmt.Println("")

	return z
}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
