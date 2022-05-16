package main

import (
	"final_project/db"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := db.OpenDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			if update.CallbackQuery.Data == "order" {
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Заполните информация о доставке в формате: /order_details Адрес, Ваш номер телефона, Ваше ФИО")
				if _, err := bot.Send(msg); err != nil {
					fmt.Println(err)
				}

				continue
			}

			cartItemID, err := strconv.Atoi(update.CallbackQuery.Data)
			if err != nil {
				fmt.Println("не могли добавить в корзину", update.CallbackQuery.Data)
			}

			err = db.InsertCartItem(cartItemID, update.CallbackQuery.Message.Chat.ID)
			if err != nil {
				fmt.Println(err)
				continue
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Позиция(%s) успешно добавлена в корзину!", update.CallbackQuery.Data))
			if _, err := bot.Send(msg); err != nil {
				fmt.Println(err)
			}

			continue
		}

		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я вас не понял")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "orders_list":
			orders, err := db.GetOrders(update.Message.Chat.ID)
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, order := range orders {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.ParseMode = "markdown"

				items, err := db.GetMenuItemsByIDs(order.ItemIdsToStringArray())
				if err != nil {
					fmt.Println(err)
					continue
				}

				for _, menuItem := range items {
					msg.Text += "Название: " + menuItem.Name + ", " + "Цена: " + strconv.Itoa(menuItem.Price) + "  "
				}

				msg.Text = strings.Trim(msg.Text, ", ")

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}

			continue
		case "menu":
			menuItems, err := db.GetMenuItems()
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, menuItem := range menuItems {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я вас не понял")
				msg.Text = fmt.Sprintf(
					`*Название*: %s 
							*Цена*: %d
							*Состав*: %s`,
					menuItem.Name, menuItem.Price, menuItem.Composition)
				msg.ParseMode = "markdown"

				addToCartButton := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", strconv.Itoa(menuItem.ID)),
					),
				)

				msg.ReplyMarkup = addToCartButton

				file := tgbotapi.FileURL(menuItem.PhotoURL)

				photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, file)

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}

				if _, err := bot.Send(photoMsg); err != nil {
					log.Panic(err)
				}
			}
			continue
		case "cart":
			menuItemsInCart, err := db.GetCartMenuItems(update.Message.Chat.ID)
			if err != nil {
				fmt.Println(err)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ParseMode = "markdown"
			msg.Text = "*В корзине у вас есть*: "

			for _, menuItem := range menuItemsInCart {
				msg.Text += menuItem.Name + ", "
			}

			msg.Text = strings.Trim(msg.Text, ", ")

			orderButton := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Сделать заказ", "order"),
				),
			)

			msg.ReplyMarkup = orderButton

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		case "order_details":
			items, err := db.GetCartMenuItems(update.Message.Chat.ID)
			if err != nil {
				fmt.Println(err)
				continue
			}

			var ids []int
			for _, item := range items {
				ids = append(ids, item.ID)
			}

			err = db.InsertOrder(db.Order{
				MenuItemIds: ids,
				CreatedAt:   time.Now(),
				ChatId:      update.Message.Chat.ID,
				Status:      "new",
				ClientInfo:  update.Message.Text,
			})
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = db.DeleteAllFromCart(update.Message.Chat.ID)
			if err != nil {
				fmt.Println(err)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Заказ успешно оформлен!")

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

			continue
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
