package main

//time for duration and sleep, fmt for printing
import (
	"fmt"
	"time"
)

/*
	This function will start a timer that will print the current time and the string (text argument) passed in.
	The pause between each printing is specified by the paus argument.
	This timer will run forever.
*/

func Remind(text string, paus time.Duration) {
	//a for loop running forever - the clock will tick forever
	for {
		time.Sleep(paus)
		fmt.Print("The time is", time.Now().Format("15:04:05"), ": ", text)
	}
}

func main() {

	// start the timers, they will run concurrently because they are goroutines
	go Remind("Time to eat\n", time.Hour*3)
	go Remind("Time to work\n", time.Hour*8)
	go Remind("Time to sleep\n", time.Hour*24)

	//don't let the program end, make it wait
	select {}

}
