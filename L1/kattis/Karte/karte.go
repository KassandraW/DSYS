package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		//map of character to map of strings to bools
		groups := map[byte]map[string]bool{
			'P': {},
			'K': {},
			'H': {},
			'T': {},
		}

		for i := 0; i < len(input); i += 3 {
			current := input[i : i+3]
			group := current[0]

			//checking for duplicates
			if groups[group][current] {
				fmt.Print("GRESKA")
				return
			}

			//set it to mark as seen?
			groups[group][current] = true
		}

		fmt.Printf("%d %d %d %d",
			13-len(groups['P']),
			13-len(groups['K']),
			13-len(groups['H']),
			13-len(groups['T']))

		/* -- My version --
		var p, k, h, t []string
		for i := 0; i < len(input); i += 3 {
			current := input[i : i+3]
			if string(current[0]) == "P" {
				if slices.Contains(p, current) {
					fmt.Print("GRESKA")
					return
				}
				p = append(p, current)

			} else if string(current[0]) == "H" {
				if slices.Contains(h, current) {
					fmt.Print("GRESKA")
					return
				}
				h = append(h, current)
			} else if string(current[0]) == "T" {
				if slices.Contains(t, current) {
					fmt.Print("GRESKA")
					return
				}
				t = append(t, current)
			} else if string(current[0]) == "K" {
				if slices.Contains(k, current) {
					fmt.Print("GRESKA")
					return
				}
				k = append(k, current)
			}
		}

		fmt.Printf("%d %d %d %d", 13-len(p), 13-len(k), 13-len(h), 13-len(t))
		*/

	}
}
