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

func (o Order) OrderStatus() string {
	statusMap := map[string]string{
		"new":       "Готовится",
		"sent":      "Отправлено курьером",
		"delivered": "Доставлено",
	}

	return statusMap[o.Status]
}

func InsertOrder(o Order) error {
	_, err := client.Exec("INSERT INTO orders (menu_item_ids, status, created_at, client_info, chat_id) VALUES ($1, $2, $3, $4, $5)",
		o.ItemIdsText(), o.Status, o.CreatedAt, o.ClientInfo, o.ChatId,
	)

	return err
}

func UpdateOrderStatus(id int, status string) error {
	_, err := client.Exec("UPDATE orders SET status = $1 WHERE id = $2",
		status, id,
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

func GetOrderByID(id int) (Order, error) {
	sqlQuery := "SELECT id, menu_item_ids, status, created_at, client_info, chat_id FROM orders WHERE id = $1"

	result, err := client.Query(sqlQuery, id)
	if err != nil {
		return Order{}, err
	}

	var orders []Order

	for result.Next() && len(orders) < 1 {
		item := Order{}
		var idsText string

		err := result.Scan(&item.Id, &idsText, &item.Status, &item.CreatedAt, &item.ClientInfo, &item.ChatId)
		if err != nil {
			fmt.Println(err)

			return Order{}, err
		}

		var menuItemIds []int

		for _, id := range strings.Split(idsText, ",") {
			menuItemId, err := strconv.Atoi(id)
			if err != nil {
				return Order{}, err
			}
			menuItemIds = append(menuItemIds, menuItemId)
		}

		item.MenuItemIds = menuItemIds

		orders = append(orders, item)
	}

	return orders[0], nil
}

func GetAllOrders() ([]Order, error) {
	sqlQuery := "SELECT id, menu_item_ids, status, created_at, client_info, chat_id FROM orders ORDER BY created_at DESC"

	result, err := client.Query(sqlQuery)
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
