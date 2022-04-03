package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
		if i == 10 {
			break
		}
	}
}
