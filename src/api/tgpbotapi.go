package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func NewBot() *tgbotapi.BotAPI {
	token, _ := os.LookupEnv("TELEGRAM_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func Update(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	return updates
}

func GetChatId(update tgbotapi.Update, bot *tgbotapi.BotAPI) int {
	if update.CallbackQuery != nil {
		return int(update.CallbackQuery.Message.Chat.ID)
	} else {
		return int(update.Message.Chat.ID)
	}
}
