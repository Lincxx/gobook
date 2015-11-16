package main

import (
	"fmt"
	"os"
)

func main() {
	//this is a comment
	//s, sep := "", ""
	for i, arg := range os.Args[0:] {
		// s += sep + arg
		// sep = " "
		fmt.Println(i, arg)
	}

}
