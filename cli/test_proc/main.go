package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Test!")
		time.Sleep(5 * time.Second)
	}
}
