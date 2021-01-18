package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/routers"
)

func Map(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery != nil {
		routers.RouteInline(update, bot)
	} else {
		routers.RouteKeyboard(update, bot)
	}
}
