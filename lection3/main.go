package main

import "fmt"

func main() {
	notebookBrand := "ACER"
	var studentsCount int
	var age int

	age = 25

	fmt.Println(notebookBrand, studentsCount, age)

	structsExamples()
}

func doSomething() {
	notebookBrand := "ACER"
	var studentsCount int
	var age int
	var acceptedToJob bool

	age = 25
	fmt.Println(notebookBrand, studentsCount, age)

	if age > 21 {
		acceptedToJob = true
		balance := 100
		fmt.Println(balance)
		fmt.Println("Ты принят на работу!")
	} else {
		acceptedToJob = false
		fmt.Println("Ты не принят на работу!")
	}

	fmt.Println(acceptedToJob)
}
