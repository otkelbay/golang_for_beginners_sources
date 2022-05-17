package main

import (
	"database/sql"
	"final_project/db"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var dbClient *sql.DB

func main() {
	err := db.OpenDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/menu", menuListHandler)
	http.HandleFunc("/orders", orderListHandler)
	http.HandleFunc("/add-menu", menuAddHandler)
	http.HandleFunc("/edit-menu", menuEditHandler)
	http.HandleFunc("/add-menu-page", menuAddPageHandler)
	http.HandleFunc("/edit-menu-page", menuEditPageHandler)
	http.HandleFunc("/delete-menu", menuDeleteHandler)
	http.HandleFunc("/change-status", changeStatusHandler)

	fmt.Println("Сервер слушает запросы по адресу: localhost:7777")

	err = http.ListenAndServe("localhost:7777", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}

func menuListHandler(w http.ResponseWriter, r *http.Request) {
	menuItems, err := db.GetMenuItems()
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
					
					<form action="http://localhost:7777/edit-menu-page">
						<input type="hidden" name="id" value="` + strconv.Itoa(menuItem.ID) + `">
						<input type="submit" class="button button1" value="Редактировать">
					</form>
				</td>
				<td>
					<form action="http://localhost:7777/delete-menu">
						<input type="hidden" name="id" value="` + strconv.Itoa(menuItem.ID) + `">
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
		
		<a href="http://localhost:7777/add-menu-page" class="button button1" style="margin-bottom: 20px;">Добавить блюдо</a>
		
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

func orderListHandler(w http.ResponseWriter, r *http.Request) {
	orderItems, err := db.GetAllOrders()
	if err != nil {
		w.Write([]byte(`Не смогли вытащить данные с БД!`))

		return
	}

	orderItemsHtml := ``

	for _, orderItem := range orderItems {
		menuItems, err := db.GetMenuItemsByIDs(orderItem.ItemIdsToStringArray())
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`Не смогли получить список!`))
		}

		totalPrice := 0
		for _, menuItem := range menuItems {
			totalPrice += menuItem.Price
		}

		compositionHtml := ``
		for _, item := range menuItems {
			compositionHtml += `
			<li>
				` + item.Name + `
			</li>	
		`
		}

		orderItemsHtml += `
			<tr>
				<td>ID</td>
				<td>
					` + orderItem.ClientInfo + `
				</td>
				<td>
					` + orderItem.CreatedAt.String() + `
				</td>
				<td>
					` + strconv.Itoa(totalPrice) + `
				</td>
				<td>
					<ul>
						` + compositionHtml + `
					</ul>
				</td>
				<td>` + orderItem.OrderStatus() + `</td>
				<td>
					<form action="http://localhost:7777/change-status" method="post">
						<input type="hidden" name="id" value="` + strconv.Itoa(orderItem.Id) + `">
						<select id="status" name="status">
							<option value="new" selected>Готовится</option>
							<option value="sent">Отправлено курьером</option>
							<option value="delivered">Доставлено</option>
						</select>
						<input type="submit" class="button button1">
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

<h2>Управление заказами</h2>

<br>

<table>
    <tr>
        <th>ID</th>
        <th>Информация о доставке</th>
        <th>Время заказа</th>
        <th>Цена заказа</th>
        <th>Состав заказа</th>
        <th>Статус заказа</th>
        <th>Поменять статус</th>
		</tr>
		` + orderItemsHtml + `
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

	position := db.MenuItem{
		Name:        r.FormValue("name"),
		PhotoURL:    r.FormValue("photo_url"),
		Price:       price,
		Composition: strings.Split(r.FormValue("composition"), ","),
	}

	_ = position

	err = db.InsertMenuItem(position)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`Не смогли добавить позицию в меню!`))

		return
	}

	w.Write([]byte(`Позиция успешно добавлена в меню!`))
}

func menuEditHandler(w http.ResponseWriter, r *http.Request) {
	idFV := r.FormValue("id")
	if idFV == "" {
		w.Write([]byte("id не должен быть пустым"))
		return
	}

	id, err := strconv.Atoi(idFV)
	if err != nil {
		w.Write([]byte("id должен быть цифрой"))
		return
	}

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

	position := db.MenuItem{
		ID:          id,
		Name:        r.FormValue("name"),
		PhotoURL:    r.FormValue("photo_url"),
		Price:       price,
		Composition: strings.Split(r.FormValue("composition"), ","),
	}

	_ = position

	err = db.UpdateMenuItem(position)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`Не смогли обновить позицию в меню!`))

		return
	}

	w.Write([]byte(`Позиция успешно обновлена в меню!`))
}

func menuAddPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<!DOCTYPE html>
		<html lang="en">
		<head>
			<title>Добавить позицию</title>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.1/dist/css/bootstrap.min.css">
			<script src="https://cdn.jsdelivr.net/npm/jquery@3.6.0/dist/jquery.slim.min.js"></script>
			<script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.1/dist/umd/popper.min.js"></script>
			<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.6.1/dist/js/bootstrap.bundle.min.js"></script>
		</head>
		<body>
		
		<div class="container">
			<h2>Добавить позицию</h2>
			<form action="http://localhost:7777/add-menu" method="post">
				<div class="form-group">
					<label for="name">Имя позиций:</label>
					<input type="text" class="form-control" id="name" placeholder="Имя позиций" name="name">
				</div>
				<div class="form-group">
					<label for="price">Цена:</label>
					<input type="number" class="form-control" id="price" placeholder="Цена" name="price">
				</div>
				<div class="form-group">
					<label for="composition">Состав:</label>
					<input type="text" class="form-control" id="composition" placeholder="Состав" name="composition">
				</div>
				<div class="form-group">
					<label for="photo_url">Ссылка на фотографию:</label>
					<input type="text" class="form-control" id="photo_url" placeholder="Ссылка" name="photo_url">
				</div>
		
				<button type="submit" class="btn btn-primary">Отправить</button>
			</form>
		</div>
		
		</body>
		</html>`))
}

func menuEditPageHandler(w http.ResponseWriter, r *http.Request) {
	idFV := r.FormValue("id")
	if idFV == "" {
		w.Write([]byte("id не должен быть пустым"))
		return
	}

	id, err := strconv.Atoi(idFV)
	if err != nil {
		w.Write([]byte("id должен быть цифрой"))
		return
	}

	menuItem, err := db.GetMenuItemByID(id)
	if err != nil {
		fmt.Println(err)

		w.Write([]byte("в базе нет такой позиций с таким id"))
		return
	}

	template := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<title>Редактировать позицию</title>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.1/dist/css/bootstrap.min.css">
		<script src="https://cdn.jsdelivr.net/npm/jquery@3.6.0/dist/jquery.slim.min.js"></script>
		<script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.1/dist/umd/popper.min.js"></script>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.6.1/dist/js/bootstrap.bundle.min.js"></script>
	</head>
	<body>
	
	<div class="container">
		<h2>Редактировать позицию: </h2>
		<form action="http://localhost:7777/edit-menu" method="post">
			<input type="hidden"  name="id" value="` + strconv.Itoa(menuItem.ID) + `">
			<div class="form-group">
				<label for="name">Имя позиций:</label>
				<input type="text" class="form-control" id="name" placeholder="Имя позиций" name="name" value="` + menuItem.Name + `">
			</div>
			<div class="form-group">
				<label for="price">Цена:</label>
				<input type="number" class="form-control" id="price" placeholder="Цена" name="price" value="` + strconv.Itoa(menuItem.Price) + `">
			</div>
			<div class="form-group">
				<label for="composition">Состав:</label>
				<input type="text" class="form-control" id="composition" placeholder="Состав" name="composition" value="` + menuItem.CompositionText() + `">
			</div>
			<div class="form-group">
				<label for="photo_url">Ссылка на фотографию:</label>
				<input type="text" class="form-control" id="photo_url" placeholder="Ссылка" name="photo_url" value="` + menuItem.PhotoURL + `">
			</div>
	
			<button type="submit" class="btn btn-primary">Отправить</button>
		</form>
	</div>
	
	</body>
	</html>
	`

	w.Write([]byte(template))
}

func menuDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("id") == "" {
		w.Write([]byte(`ID не должен быть пустым!`))

		return
	}

	menuItemIDString := r.FormValue("id")
	menuItemID, err := strconv.Atoi(menuItemIDString)
	if err != nil {
		w.Write([]byte(`ID должен быть цифрой!`))

		return
	}

	err = db.DeleteMenuItem(menuItemID)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`Не удалось удалить!`))

		return
	}

	w.Write([]byte(`Позиция успешно удалена из меню!`))
}

func changeStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("id") == "" {
		w.Write([]byte(`ID не должен быть пустым!`))

		return
	}

	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.Write([]byte(`ID должен быть цифрой!`))

		return
	}

	if r.FormValue("status") == "" {
		w.Write([]byte(`статус не должен быть пустым!`))

		return
	}

	order, err := db.GetOrderByID(id)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`Не наши такого заказа!`))

		return
	}

	err = db.UpdateOrderStatus(id, r.FormValue("status"))
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`Не удалось обновить статус заказа!`))

		return
	}

	SendMessageToClient(order.ChatId, fmt.Sprintf("Новый статус заказа: %s", order.OrderStatus()))

	w.Write([]byte(`Статус заказа успешно обновлен!`))
}

var botAPI *tgbotapi.BotAPI

func SendMessageToClient(chatId int64, message string) {
	var err error
	if botAPI == nil {
		botAPI, err = tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
		if err != nil {
			log.Panic(err)
		}

		botAPI.Debug = true
	}

	msg := tgbotapi.NewMessage(chatId, message)

	if _, err := botAPI.Send(msg); err != nil {
		fmt.Println(err)
	}
}