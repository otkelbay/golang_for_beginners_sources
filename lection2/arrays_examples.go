package main

import "fmt"

func testArrays() {

	names := []string{"Nurlan", "Yerlan"}
	names = append(names, "Nursultan")
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(names[0])
	fmt.Println(names[0:2])
	names[1] = "Arman"
	fmt.Println(names)

}
