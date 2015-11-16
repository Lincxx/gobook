package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	stop := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "), start, stop)
}
