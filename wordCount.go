/*
	@author - Ivan Liljeqvist
	@version - 23-03-2015
*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

/*
	This program includes a function that counts the number of times a string contains each word.
*/

/*
	This function counts the number of times the string that is passed in as the parameter contains each word.
	The results are returned as a map.
*/

func WordCount(s string) map[string]int {

	//extract words from the string
	words := strings.Fields(s)
	//initialize the array that will be returned
	wordCount := make(map[string]int)

	//loop through the words, increase the count of words
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
