/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-08
 * @fileoverview Is the word in the given sentence
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// declare variables
	var sentence string
	var word string
	var exist bool = false
	reader := bufio.NewReader(os.Stdin)

	// input
	// sentence
	fmt.Print("Please enter a sentence: ")
	sentence, _ = reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	// word to be searched
	fmt.Print("Enter a word to search in your sentence: ")
	word, _ = reader.ReadString('\n')
	word = strings.TrimSpace(word)

	// find if the word is in the sentence 
	for counter := 0; counter <= len(sentence)-len(word); counter++ {

		var match bool = true

		for notFound := 0; notFound < len(word); notFound++ {
			if sentence[counter+notFound] != word[notFound] {
				match = false
			}
		}

		if match {
			exist = true
		}
	}

	// output
	if exist {
		fmt.Printf("Hooray, the word %s was found in the sentence: %s", word, sentence)
	} else {
		fmt.Printf("Sorry, the word %s was not found in the sentence: %s", word, sentence)
	}

	fmt.Println("\nDone.")
}
