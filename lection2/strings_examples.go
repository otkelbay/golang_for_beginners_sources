package main

import (
	"fmt"
	"strings"
)

func testString1() {

	someRandomText := "Kayrat won its first game in a season."
	fcName := someRandomText[0:6]

	someRandomTextInRussian := "Кайрат выиграл свой первый матч в сезоне."
	fcNameRussian := someRandomTextInRussian[0:6]
	fcNameRussianFull := someRandomTextInRussian[0:12]

	fmt.Println(fcName)
	fmt.Println(fcNameRussian)
	fmt.Println(fcNameRussianFull)
	fmt.Println(someRandomText[:01]) //Получаем одну букву

	/*
		Kayrat
		Кай
		Кайрат
		K
	*/

	containsKayrat := strings.Contains(someRandomText, "Kayrat")
	containsKaysar := strings.Contains(someRandomText, "Kaysar")
	gamePos := strings.Index(someRandomText, "game")
	lower := strings.ToLower("OtKeLBaY")
	replaceToKaysar := strings.Replace(someRandomText, "Kayrat", "Kaysar", 1)

	fmt.Println(containsKayrat)
	fmt.Println(containsKaysar)
	fmt.Println(gamePos)
	fmt.Println(lower)
	fmt.Println(replaceToKaysar)

	/*
		true
		false
		21
		otkelbay
		Kaysar won its first game in a season.
	*/

	for i, letter := range fcName {
		fmt.Println(i, string(letter))
	}
}
