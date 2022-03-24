package main

import "fmt"

func testMap() {

	var nameAgeMap map[string]int
	nameAgeMap = make(map[string]int)

	nameAgeMap["Nurlan"] = 25
	nameAgeMap["Yerlan"] = 30
	nameAgeMap["Petya"] = 32

	petyaAge := nameAgeMap["Petya"]

	fmt.Println(petyaAge)
}
