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

	nurlan := person{"Nurlan", 23}

	vanya := person{name: "Vanya"}

	secretSpy := person{}

	fmt.Println(tom, nurlan, vanya, secretSpy)

	fmt.Println(tom.age, tom.name)

	tom.age = 40
	tom.name = "Jerry"

	fmt.Println(tom)

	tom.changeNameAndAge("Danny", 22)
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
