package main

import (
	"fmt"
	"strconv"
	"strings"
)

type person struct {
	name string
	age  int
}

func structsExamples() {
	var tom person
	tom = person{"Tom", 24}

	tom.age = 40
	tom.name = "Jerry"

	fmt.Println(tom.age)

	tom.changeNameAndAge("Danny", 22)

	fmt.Println(tom)
	nickName := tom.generateNickName()
	tom.eat("бургер")

	fmt.Println(nickName)
}

func (p *person) changeNameAndAge(newName string, newAge int) {
	p.name = newName
	p.age = newAge
	fmt.Println(p)
}

func (p person) generateNickName() string {
	return strings.ToLower(p.name) + "#" + strconv.Itoa(p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}
