package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
	"strings"
)

type MenuItem struct {
	ID          int
	Price       int
	PhotoURL    string
	Composition []string
	Name        string
}

var cartItems []int

func main() {
	menuItems := []MenuItem{
		{
			ID:          1,
			Price:       3000,
			PhotoURL:    "https://chefrestoran.ru/wp-content/uploads/2018/10/604655519.jpg",
			Composition: []string{"Кола", "Фри", "Бургер"},
			Name:        "Комбо чизбургер",
		},
		{
			ID:          2,
			Price:       2000,
			PhotoURL:    "https://static.1000.menu/img/content-v2/05/d8/21554/klassicheskaya-shaurma_1589963797_11_max.jpg",
			Composition: []string{"Кола", "Фри", "Шаурма"},
			Name:        "Комбо шаурма",
		},
		{
			ID:          3,
			Price:       1000,
			PhotoURL:    "https://www.maggi.ru/data/images/recept/img640x500/recept_5828_bnu7.jpg",
			Composition: []string{"Кола", "Фри", "Самса"},
			Name:        "Комбо",
		},
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
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.

			fmt.Println(update.CallbackQuery.Data)

			cartItemID, err := strconv.Atoi(update.CallbackQuery.Data)
			if err != nil {
				fmt.Println("не могли добавить в корзину", update.CallbackQuery.Data)
			}

			cartItems = append(cartItems, cartItemID)
			fmt.Println(cartItems)

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Позиция(%s) успешно добавлена в корзину!", update.CallbackQuery.Data))
			if _, err := bot.Send(msg); err != nil {
				panic(err)
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
			msg.Text = "Функция списка заказов будет чуть позже!"
		case "menu":
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
			var menuItemsInCart []MenuItem
			for _, menuItem := range menuItems {
				for _, cartID := range cartItems {
					if menuItem.ID == cartID {
						menuItemsInCart = append(menuItemsInCart, menuItem)
					}
				}
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ParseMode = "markdown"
			msg.Text = "*В корзине у вас есть*: "

			for _, menuItem := range menuItemsInCart {
				msg.Text += menuItem.Name + ", "
			}

			msg.Text = strings.Trim(msg.Text, ", ")

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		case "order_details":
			fmt.Println(update.Message.Text)
			continue
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
