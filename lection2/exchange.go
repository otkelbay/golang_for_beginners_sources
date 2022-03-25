package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exchange() {
	var currenciesBalanceMap map[string]float64
	var currenciesExchangeBuyingRate map[string]float64
	var currenciesExchangeSellingRate map[string]float64

	currenciesBalanceMap = make(map[string]float64)
	currenciesExchangeBuyingRate = make(map[string]float64)
	currenciesExchangeSellingRate = make(map[string]float64)

	currenciesBalanceMap["EUR"] = 300
	currenciesBalanceMap["KZT"] = 100000
	currenciesBalanceMap["USD"] = 1000

	currenciesExchangeSellingRate["EUR"] = 515
	currenciesExchangeSellingRate["USD"] = 415

	currenciesExchangeBuyingRate["EUR"] = 500
	currenciesExchangeBuyingRate["USD"] = 400

	fmt.Println(currenciesBalanceMap)
	fmt.Println("Курс на покупку", currenciesExchangeBuyingRate)
	fmt.Println("Курс на продажу", currenciesExchangeSellingRate)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if text == "EXIT" {
			return
		}

		operands := strings.Split(text, " ")
		fmt.Println(operands)
		fmt.Println(text)

		operation := operands[0]
		currency := operands[1]
		amount, err := strconv.ParseFloat(operands[2], 64)
		if err != nil {
			fmt.Println("Количество должно быть цифрой")
			continue
		}

		if operation == "BUY" {
			leftBalance := currenciesBalanceMap[currency]
			if leftBalance >= amount {
				currenciesBalanceMap[currency] = currenciesBalanceMap[currency] - amount
				currenciesBalanceMap["KZT"] = currenciesBalanceMap["KZT"] + currenciesExchangeSellingRate[currency]*amount
				fmt.Println("Вы успешно купили", currency)
			} else {
				fmt.Println("У нас нет столько то ", currency)
			}
		} else if operation == "SELL" {
			leftBalance := currenciesBalanceMap["KZT"]
			if leftBalance >= currenciesExchangeBuyingRate[currency]*amount {
				currenciesBalanceMap["KZT"] = currenciesBalanceMap["KZT"] - currenciesExchangeBuyingRate[currency]*amount
				currenciesBalanceMap[currency] = currenciesBalanceMap[currency] + amount
				fmt.Println("Вы успешно продали", currency)
			} else {
				fmt.Println("У нас нет столько денег в тенге чтобы купить у вас", currency)
			}

		} else {
			fmt.Println("Операция должна быть либо BUY либо SELL")
		}

		fmt.Println(currenciesBalanceMap)

		continue
	}

}
