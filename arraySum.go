/*
	@author - Ivan Liljeqvist
	@version - 23-03-2015
*/

package main

/*
	This program adds the elements in an array by using two go routines.
	The first routine adds the first half of the array and the second routine adds the second half of the routine.
	When this is done the sum is printed to the console.
*/

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {

	sum := 0

	//go through all numbers in array and add them to the sum
	for _, i := range a {
		sum += i
	}

	//write the result to the channel
	res <- sum
}

func main() {
	//init the array
	a := []int{1, 2, 3, 4, 5, 6, 7}

	n := len(a)
	ch := make(chan int)

	//start the routines and give them half of the array each
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	//wait and save the first answer
	firstHalf := <-ch
	//wait and save the second answer
	secondHalf := <-ch

	//add the answers together
	sum := firstHalf + secondHalf

	println(sum)
}
