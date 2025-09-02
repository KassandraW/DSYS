package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	//first line
	s.Scan()

	//second line
	s.Scan()

	words := strings.Fields(s.Text())

	conclusion := "makes sense"

	for i, word := range words {
		if word == "mumble" {
			continue
		} else if r, err := strconv.Atoi(word); err == nil {
			if r != i+1 {
				conclusion = "something is fishy"
				break
			}
		}
	}

	fmt.Print(conclusion)
}
