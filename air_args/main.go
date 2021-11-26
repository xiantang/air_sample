package main

import (
	"fmt"
	"os"
)

func main() {
	// get and print env variables
	fmt.Println("ENV:", os.Getenv("ENV"))
	// get and print os second argument
	fmt.Println("ARG:", os.Args[1])

}
