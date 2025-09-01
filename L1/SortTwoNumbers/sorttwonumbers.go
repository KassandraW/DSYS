package main

import (
	"fmt"
)

func main() { 
	var A, B int 

	fmt.Scan(&A, &B)

	if A < B {
		fmt.Printf( "%d %d", A, B )
	} else {
		fmt.Printf( "%d %d", B, A )
	}
	
}