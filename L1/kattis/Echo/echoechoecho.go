package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() { 
	Scanner := bufio.NewScanner(os.Stdin)
	if Scanner.Scan() {
		input := Scanner.Text()
		fmt.Print(input + " " + input + " " + input )
	}

	/*var word string; 
    fmt.Scan(&word)
    fmt.Print(word + " " + word + " " + word ) */
	
}