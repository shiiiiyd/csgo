package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello, 世界！", os.Args[1])
	}
}
