package main

import "fmt"

func main() {
	notebookBrand := "ACER"
	var studentsCount int
	var age int

	age = 25

	fmt.Println(notebookBrand, studentsCount, age)

	structsExamples()

	//0 - нолики
	//1 - X
	//2 - пустое
	var krestikiNoliki [][]int
	krestikiNoliki = [][]int{
		{2, 2, 1},
		{2, 0, 2},
		{10, 2, 2},
	}

	fmt.Println(krestikiNoliki[2][0])
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
		if acceptedToJob {
			fmt.Println(balance)
		}
	} else {
		acceptedToJob = false
		fmt.Println("Ты не принят на работу!")
	}

	fmt.Println(acceptedToJob)
}
