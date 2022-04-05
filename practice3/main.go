package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 0 - нолик, 1 - крестик, 2 - пусто
	var xesAndZeroes [][]int
	xesAndZeroes = [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}

	// false - это ход нолика, true - это ход крестика
	turnOfX := false

	reader := bufio.NewReader(os.Stdin)

	for {
		displayGameMap(xesAndZeroes)
		fmt.Print("-> Ход ")
		if turnOfX {
			fmt.Print("крестиков: ")
		} else {
			fmt.Print("ноликов: ")
		}
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)

		if text == "EXIT" {
			break
		}

		positionEls := strings.Split(text, ",")
		if len(positionEls) != 2 {
			fmt.Println("Введите позицию в формате строка,столбец например: 2,1")
			continue
		}

		row, err := strconv.Atoi(positionEls[0])
		if err != nil {
			fmt.Println("Введите позицию в формате строка,столбец например: 2,1")
			continue
		}

		column, err := strconv.Atoi(positionEls[1])
		if err != nil {
			fmt.Println("Введите позицию в формате строка,столбец например: 2,1")
			continue
		}

		if row > 2 || column > 2 {
			fmt.Println("Максимальное значение позиции это 2")
			continue
		}

		if xesAndZeroes[row][column] != 2 {
			fmt.Println("Ячейка занята.")
			continue
		}

		if turnOfX {
			xesAndZeroes[row][column] = 1
		} else {
			xesAndZeroes[row][column] = 0
		}

		winner := defineIsGameHasWinner(xesAndZeroes)
		if winner == 2 {
			turnOfX = !turnOfX
			continue
		}

		displayGameMap(xesAndZeroes)

		if winner == 1 {
			fmt.Println("Крестик выиграл!")
		} else if winner == 0 {
			fmt.Println("Нолик выиграл!")
		}

		break
	}

}

func displayGameMap(state [][]int) {
	fmt.Println("Поле:")
	fmt.Println("-------")
	for _, row := range state {
		items := []string{}
		for _, el := range row {
			if el == 2 {
				items = append(items, " ")
			} else {
				items = append(items, strconv.Itoa(el))
			}
		}
		fmt.Println("|" + strings.Join(items, "|") + "|")

		fmt.Println("-------")
	}
	fmt.Println()
}

// 0 - нолик, 1 - крестик, 2 - победителя пока нет
func defineIsGameHasWinner(state [][]int) int {
	if state[0][0] == 0 && state[0][1] == 0 && state[0][2] == 0 {
		return 0
	} else if state[1][0] == 0 && state[1][1] == 0 && state[1][2] == 0 {
		return 0
	} else if state[2][0] == 0 && state[2][1] == 0 && state[2][2] == 0 {
		return 0
	} else if state[0][0] == 0 && state[1][0] == 0 && state[2][0] == 0 {
		return 0
	} else if state[0][1] == 0 && state[1][1] == 0 && state[2][1] == 0 {
		return 0
	} else if state[0][2] == 0 && state[1][2] == 0 && state[2][2] == 0 {
		return 0
	} else if state[0][0] == 0 && state[1][1] == 0 && state[2][2] == 0 {
		return 0
	} else if state[0][2] == 0 && state[1][1] == 0 && state[2][0] == 0 {
		return 0
	}

	if state[0][0] == 1 && state[0][1] == 1 && state[0][2] == 1 {
		return 1
	} else if state[1][0] == 1 && state[1][1] == 1 && state[1][2] == 1 {
		return 1
	} else if state[2][0] == 1 && state[2][1] == 1 && state[2][2] == 1 {
		return 1
	} else if state[0][0] == 1 && state[1][0] == 1 && state[2][0] == 1 {
		return 1
	} else if state[0][1] == 1 && state[1][1] == 1 && state[2][1] == 1 {
		return 1
	} else if state[0][2] == 1 && state[1][2] == 1 && state[2][2] == 1 {
		return 1
	} else if state[0][0] == 1 && state[1][1] == 1 && state[2][2] == 1 {
		return 1
	} else if state[0][2] == 1 && state[1][1] == 1 && state[2][0] == 1 {
		return 1
	}

	return 2
}
