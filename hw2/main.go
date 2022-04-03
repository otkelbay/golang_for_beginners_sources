package main

import (
	"fmt"
	"strconv"
	"strings"
)

func filterStudentsToClasses() {
	var names []string
	names = []string{
		"Саша", "Маша", "Ерлан", "Арман", "Ян", "Юля", "Данияр", "Константин", "Михаил", "Василий",
	}

	var studentsA []string
	var studentsB []string
	var studentsC []string

	for _, name := range names {
		firstLetter := name[0:2]
		if firstLetter == "А" || firstLetter == "Б" || firstLetter == "В" || firstLetter == "Г" || firstLetter == "Д" || firstLetter == "Е" {
			studentsA = append(studentsA, name)
		} else if firstLetter == "Ж" || firstLetter == "З" || firstLetter == "К" ||
			firstLetter == "Л" || firstLetter == "М" || firstLetter == "Н" ||
			firstLetter == "И" || firstLetter == "О" || firstLetter == "Р" {
			studentsB = append(studentsB, name)
		} else {
			studentsC = append(studentsC, name)
		}
	}

	fmt.Println("А: " + strings.Join(studentsA, ",") + ". " + strconv.Itoa(len(studentsA)) + " ученика.")
	fmt.Println("B: " + strings.Join(studentsB, ",") + ". " + strconv.Itoa(len(studentsB)) + " ученика.")
	fmt.Println("C: " + strings.Join(studentsC, ",") + ". " + strconv.Itoa(len(studentsC)) + " ученика.")

}

func main() {
	filterStudentsToClasses()
}
