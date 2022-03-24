package main

import (
	"fmt"
	"strconv"
)

func testConvertions() {

	balance := float64(456.0)
	fmt.Println(int(balance))
	fmt.Println(int('A'))
	fmt.Println(float64('A'))
	//fmt.Println(float64("A"))
	fmt.Println(float64(-555324))

	fmt.Println(string(1041))
	fmt.Println(strconv.Itoa(1041))
	fmt.Println(strconv.Atoi("1041"))
	fmt.Println(strconv.ParseBool("false"))

	someInt := 50
	someFloat := 111.1

	result := float64(someInt) + someFloat

	fmt.Println(result)
}
