package routers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/controllers"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/message"
)

func RouteKeyboard(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	switch update.Message.Text {
	case "Каталог":
		message.Send(update, bot, message.Get(update.Message.Text))
		controllers.GetCatalog(update, bot)

	case "Корзина":
		message.Send(update, bot, message.Get(update.Message.Text))
		controllers.GetCart(update, bot)

	case "Заказы":
		controllers.ShowOrders(update, bot)

	case "Оформить заказ":
		controllers.Checkout(update, bot)
	}
}
