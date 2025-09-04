package main

import (
	"fmt"
)


func A(ch1, ch2 chan string){
	ch1 <- "ping"
	x := <- ch2
	fmt.Println(x)
}

func B(ch1, ch2 chan string){
	x := <- ch1 
	fmt.Println(x)
	ch2 <- "pong"
}
func main (){
	var ch1 = make(chan string, 4)
	var ch2 = make(chan string, 4)
	for {
		go A(ch1,ch2)
		go B(ch1, ch2)
	}
}

//Exercise 1.4
/* My channels are synchronous (unbuffered). If they had a buffer
size greater than 0, there might be stored lots pings in ch1 */