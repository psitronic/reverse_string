package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)


func reverse(text string) string {
/* The function reverses a text
   Takes one string argument
   Returns string
*/	
	//Convert string into slice and strip "\n" at the end
	//Should be improved and platform independent
	splitText := strings.Split(text, "")[:len(text)-1]
	
	//Swap characters
	// Iterates only up to the middle of the text
	//if there is an odd number of characters, the mid char should not be swapped
	for i := 0; i < len(splitText)/2; i++ {
		splitText[i], splitText[len(splitText) - i - 1] = splitText[len(splitText) - i - 1], splitText[i]
	}
	return strings.Join(splitText,"")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Printf("Reversed text: %s\n", reverse(text))
}
