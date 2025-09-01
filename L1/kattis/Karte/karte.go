package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() { 
	Scanner := bufio.NewScanner(os.Stdin)
	if Scanner.Scan() {
		input := Scanner.Text()

		for i, v := range input[:3] {
			fmt.Printf("2**%d = %d\n", i, v)
		}
	}
}