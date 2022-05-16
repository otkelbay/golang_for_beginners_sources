package db

import (
	"fmt"
	"strings"
)

type MenuItem struct {
	ID          int
	Price       int
	PhotoURL    string
	Composition []string
	Name        string
}

func (m MenuItem) CompositionText() string {
	return strings.Join(m.Composition, ",")
}

func GetMenuItems() ([]MenuItem, error) {
	sqlQuery := "SELECT id, price, photo_url, composition, name FROM menu_items"

	result, err := client.Query(sqlQuery)
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

func InsertMenuItem(menuItem MenuItem) error {
	_, err := client.Exec("INSERT INTO menu_items (price, photo_url, composition, name) VALUES ($1, $2, $3, $4)",
		menuItem.Price, menuItem.PhotoURL, menuItem.CompositionText(), menuItem.Name,
	)

	return err
}
