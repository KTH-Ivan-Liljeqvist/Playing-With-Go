/*
	@author - Ivan Liljeqvist
	@version - 23-03-2015
*/

package main

import "fmt"
import "math"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	//initialize last and last-last values
	lastValue := 0
	lastLastValue := -1

	//the closure to return
	nextFibNumb := func() int {

		//this will give us the next value, but it will be negative
		nextValueButNegative := lastValue + lastLastValue
		//math.Abs will return a float64 value
		nextValueFloat := math.Abs(float64(nextValueButNegative))
		//the function has to return int - convert to int
		nextValueInt := int(nextValueFloat)

		//update these
		lastLastValue = lastValue
		lastValue = nextValueInt

		return nextValueInt
	}

	return nextFibNumb
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
