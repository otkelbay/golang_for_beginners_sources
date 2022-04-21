package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var dbClient *sql.DB

func main() {
	var err error

	dbClient, err = openDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/menu", menuListHandler)
	http.HandleFunc("/add-menu", menuAddHandler)
	http.HandleFunc("/delete-menu", menuDeleteHandler)

	fmt.Println("Сервер слушает запросы по адресу: localhost:7777")

	err = http.ListenAndServe("localhost:7777", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}

func menuListHandler(w http.ResponseWriter, r *http.Request) {
	priceParam := r.FormValue("price")
	price, err := strconv.Atoi(priceParam)

	menuItems, err := GetItemsFromDB(price)
	if err != nil {
		w.Write([]byte(`Не смогли вытащить данные с БД!`))

		return
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
					<form action="http://localhost:7777/delete-menu">
						<input type="hidden" name="name" value="` + menuItem.Name + `">
						<input type="submit" class="button button2" value="Удалить">
					</form>
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

func menuAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("name") == "" {
		w.Write([]byte(`Название не должно быть пустым!`))

		return
	}

	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		w.Write([]byte(`Цена должна быть правильным целым числом!`))

		return
	}

	if price <= 0 {
		w.Write([]byte(`Цена должна быть выше 0!`))

		return
	}

	if r.FormValue("photo_url") == "" {
		w.Write([]byte(`Ссылка на фото не должна быть пустой!`))

		return
	}

	if r.FormValue("composition") == "" {
		w.Write([]byte(`Состав не должен быть пустым!`))

		return
	}

	strings.Split(r.FormValue("composition"), ",")

	position := MenuItem{
		Name:        r.FormValue("name"),
		PhotoURL:    r.FormValue("photo_url"),
		Price:       price,
		Composition: strings.Split(r.FormValue("composition"), ","),
	}

	_ = position

	err = InsertMenuItemToDB(position)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`Не смогли добавить позицию в меню!`))

		return
	}

	w.Write([]byte(`Позиция успешно добавлена в меню!`))
}

func menuDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("name") == "" {
		w.Write([]byte(`Название не должно быть пустым!`))

		return
	}

	deleteName := r.FormValue("name")

	_ = deleteName

	w.Write([]byte(`Позиция успешно удалена из меню!`))
}
