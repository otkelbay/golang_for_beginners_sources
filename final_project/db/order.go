package db

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Order struct {
	Id          int
	MenuItemIds []int
	Status      string
	CreatedAt   time.Time
	ClientInfo  string
	ChatId      int64
}

func InsertOrder(o Order) error {
	_, err := client.Exec("INSERT INTO orders (menu_item_ids, status, created_at, client_info, chat_id) VALUES ($1, $2, $3, $4, $5)",
		o.ItemIdsText(), o.Status, o.CreatedAt, o.ClientInfo, o.ChatId,
	)

	return err
}

func (o Order) ItemIdsText() string {
	var ids []string
	for _, mid := range o.MenuItemIds {
		ids = append(ids, strconv.Itoa(mid))
	}

	return strings.Join(ids, ",")
}

func (o Order) ItemIdsToStringArray() []string {
	var ids []string
	for _, mid := range o.MenuItemIds {
		ids = append(ids, strconv.Itoa(mid))
	}

	return ids
}

func GetOrders(chatId int64) ([]Order, error) {
	sqlQuery := "SELECT id, menu_item_ids, status, created_at, client_info, chat_id FROM orders WHERE chat_id = $1"

	result, err := client.Query(sqlQuery, chatId)
	if err != nil {
		return nil, err
	}

	var orders []Order

	for result.Next() {
		item := Order{}
		var idsText string

		err := result.Scan(&item.Id, &idsText, &item.Status, &item.CreatedAt, &item.ClientInfo, &item.ChatId)
		if err != nil {
			fmt.Println(err)

			return nil, err
		}

		var menuItemIds []int

		for _, id := range strings.Split(idsText, ",") {
			menuItemId, err := strconv.Atoi(id)
			if err != nil {
				return nil, err
			}
			menuItemIds = append(menuItemIds, menuItemId)
		}

		item.MenuItemIds = menuItemIds

		orders = append(orders, item)
	}

	return orders, nil
}
