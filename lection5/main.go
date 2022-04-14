package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type MenuItem struct {
	Price       int
	PhotoURL    string
	Composition []string
	Name        string
}

var menuItems []MenuItem

func main() {
	menuItems = []MenuItem{
		{
			Price:       3000,
			PhotoURL:    "https://chefrestoran.ru/wp-content/uploads/2018/10/604655519.jpg",
			Composition: []string{"Кола", "Фри", "Бургер"},
			Name:        "Комбо чизбургер",
		},
		{
			Price:       2000,
			PhotoURL:    "https://static.1000.menu/img/content-v2/05/d8/21554/klassicheskaya-shaurma_1589963797_11_max.jpg",
			Composition: []string{"Кола", "Фри", "Шаурма"},
			Name:        "Комбо шаурма",
		},
		{
			Price:       1000,
			PhotoURL:    "https://www.maggi.ru/data/images/recept/img640x500/recept_5828_bnu7.jpg",
			Composition: []string{"Кола", "Фри", "Самса"},
			Name:        "Комбо",
		},
	}

	http.HandleFunc("/menu", menuListHandler)

	err := http.ListenAndServe("localhost:7777", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}

func menuListHandler(w http.ResponseWriter, r *http.Request) {
	priceParam := r.FormValue("price")
	price, err := strconv.Atoi(priceParam)
	if err == nil && price != 0 {
		var filteredMenu []MenuItem
		for _, menuItem := range menuItems {
			if menuItem.Price <= price {
				filteredMenu = append(filteredMenu, menuItem)
			}
		}
		menuItems = filteredMenu
	}

	menuItemsHtml := ``

	for _, menuItem := range menuItems {
		compositionHtml := ``
		for _, item := range menuItem.Composition {
			compositionHtml += `
			<li>
				` + item + `
			</li>	
		`
		}

		menuItemsHtml += `
			<tr>
				<td>` + menuItem.Name + `</td>
				<td>
					<img src="` + menuItem.PhotoURL + `" alt="Бургер"
						 class="menu-item-photo">
				</td>
				<td>
					<ul>
						` + compositionHtml + `
					</ul>
				</td>
				<td>` + strconv.Itoa(menuItem.Price) + ` тенге</td>
				<td>
					<button class="button button1">Редактировать</button>
				</td>
				<td>
					<button class="button button2">Удалить</button>
				</td>
			</tr>
		`
	}

	template := `
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				table {
					font-family: arial, sans-serif;
					border-collapse: collapse;
					width: 100%;
				}
		
				td, th {
					border: 1px solid #dddddd;
					text-align: left;
					padding: 8px;
				}
		
				tr:nth-child(even) {
					background-color: #dddddd;
				}
		
				.menu-item-photo {
					height: 100px;
					width: auto;
				}
				.button {
					border: none;
					color: white;
					padding: 15px 32px;
					text-align: center;
					text-decoration: none;
					display: inline-block;
					font-size: 16px;
					margin: 4px 2px;
					cursor: pointer;
				}
		
				.button1 {background-color: #4ca0af;} /* Green */
				.button2 {background-color: #ba0016;} /* Blue */
			</style>
			<meta charset="utf-8">
		</head>
		<body>
		
		<h2>Управление меню</h2>

		<form action="http://localhost:7777/menu">
			<label for="price">Цена до:</label><br>
			<input type="text" id="price" name="price"><br>
			<input type="submit" value="Фильтр">
		</form>
		
		<br>
		
		<table>
			<tr>
				<th>Название</th>
				<th>Фото</th>
				<th>Состав</th>
				<th>Цена</th>
				<th>Редактировать</th>
				<th>Удалить</th>
			</tr>
			` + menuItemsHtml + `
		</table>
		
		</body>
		</html>
		`

	w.Write([]byte(template))
}

func uploadFileImage(w http.ResponseWriter, r *http.Request) {
	image, _, err := r.FormFile("item_image")
	if err != nil {
		log.Fatal(err)
	}

	defer image.Close()

	fileBytes, err := ioutil.ReadAll(image)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("image.png", fileBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
