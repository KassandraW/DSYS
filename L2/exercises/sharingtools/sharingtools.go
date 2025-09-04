package main

import (
	"fmt"
	"time"
)

func Manny(hammer, nailbox chan string) {
	for {
		<- hammer
		<- nailbox 
		fmt.Println("Manny is building");
		time.Sleep( 500 * time.Millisecond)
		hammer <- "hammer"
		fmt.Println("Manny is resting");
		time.Sleep( 500 * time.Millisecond)
		nailbox <- "nail"
	}
}

func Bob(hammer, nailbox chan string ) {
	for {
		x := <- hammer
		y := <- nailbox
		fmt.Println("Bob is building")
		time.Sleep( 500 * time.Millisecond)
		hammer <- x
		fmt.Println("Bob is resting")
		time.Sleep( 500 * time.Millisecond)
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

	time.Sleep(5 * time.Second)
}