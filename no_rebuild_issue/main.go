package main

import (
	"fmt"
	"time"
)

func main() {

	for {

		fmt.Printf("a1: %s | b1: %s\n", string(1), string(11))
		time.Sleep(time.Second * 5)
	}

}
