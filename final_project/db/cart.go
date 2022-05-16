package db

import (
	"fmt"
	"strconv"
	"strings"
)

func GetCartMenuItems(chatId int64) ([]MenuItem, error) {
	sqlQuery := "SELECT menu_item_id FROM cart where chat_id = $1"

	result, err := client.Query(sqlQuery, chatId)
	if err != nil {
		return nil, err
	}

	var menuItemIDs []string

	for result.Next() {
		var id int
		err := result.Scan(&id)
		if err != nil {
			fmt.Println(err)

			return nil, err
		}

		menuItemIDs = append(menuItemIDs, strconv.Itoa(id))
	}

	return GetMenuItemsByIDs(menuItemIDs)
}

func InsertCartItem(menuItemId int, chatId int64) error {
	_, err := client.Exec("INSERT INTO cart (menu_item_id, chat_id) VALUES ($1, $2)",
		menuItemId, chatId,
	)

	return err
}

func DeleteAllFromCart(chatId int64) error {
	_, err := client.Exec("DELETE FROM cart WHERE chat_id = $1",
		chatId,
	)

	return err
}

func GetMenuItemsByIDs(menuItemIDs []string) ([]MenuItem, error) {
	cartItemsSqlQuery := "SELECT id, price, photo_url, composition, name FROM menu_items WHERE id in (" + strings.Join(menuItemIDs, ",") + ")"

	result, err := client.Query(cartItemsSqlQuery)
	if err != nil {
		return nil, err
	}

	var menuItems []MenuItem

	for result.Next() {
		item := MenuItem{}
		var compositionText string

		err := result.Scan(&item.ID, &item.Price, &item.PhotoURL, &compositionText, &item.Name)
		if err != nil {
			fmt.Println(err)

			return nil, err
		}

		item.Composition = strings.Split(compositionText, ",")

		menuItems = append(menuItems, item)
	}

	return menuItems, nil
}
