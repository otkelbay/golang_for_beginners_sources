package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MenuItem struct {
	ID          int
	Price       int
	PhotoURL    string
	Composition []string
	Name        string
}

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
				msg.Text = fmt.Sprintf(`Название: %s '\n' Цена: %d
				Состав: %s`, menuItem.Name, menuItem.Price, menuItem.Composition)
				msg.ParseMode = "markdown"

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}
			continue
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
