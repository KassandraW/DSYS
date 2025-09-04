package main

import (
	"fmt"
)

func Manny(hammer, nailbox chan string) {
	for {
		<- hammer
		<- nailbox 
		fmt.Println("Manny is building");
		hammer <- "hammer"
		fmt.Println("Manny is resting");
		nailbox <- "nail"
	}
}

func Bob(hammer, nailbox chan string ) {
	for {
		x := <- hammer
		y := <- nailbox
		fmt.Println("Bob is building")
		hammer <- x
		fmt.Println("Bob is resting")
		nailbox <- y
	}
}


func main(){
	var hammer = make(chan string)
	var nailbox = make(chan string)
	
	go Manny(hammer, nailbox)
	go Bob(hammer, nailbox)
	hammer <- "hammer"
	nailbox <- "nail"

	for{}
	
}