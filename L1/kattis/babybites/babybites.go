package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	//to handle the next line with spaces
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') //discard leftover newline?

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s) // remove the trailing newline?
	fmt.Println(s)

	//for i = 0; 9 <

	//conclusion := "makes sense"

	/*String conclusion = "makes sense";


	  for(int i = 1; i <= n ; i++){
	      String iS = Integer.toString(i);
	      String t = s.next();
	      if(t.equals("mumble")){
	      } else if (t.equals(iS) != true){
	      	conclusion = "Something is fishy";
	       }
	  }
	      System.out.println(conclusion); */
}
