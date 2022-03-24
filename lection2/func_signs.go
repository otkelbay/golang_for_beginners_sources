package main

import "fmt"

func doSomething() {
	fmt.Println("2+2=4")
}

func printNames(names []string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getCalculationSum() int {
	return 2 + 2
}

func getSumOfNums(a, b int) int {
	return a + b
}

func getInitials(name, surname string) string {
	return name[0:1] + "." + surname[0:1] + "."
}
