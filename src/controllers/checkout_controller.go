package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/checkout"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/domain/clients"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/domain/orders"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/message"
	"strconv"
)

func Handler(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if checkout.IsName == true && update.Message != nil {
		checkout.Name = update.Message.Text
		checkout.GetPhone(update, bot)
		return
	}

	if checkout.IsPhone == true && update.Message != nil {
		checkout.Phone = update.Message.Text
		checkout.GetAddress(update, bot)
		return
	}

	if checkout.IsAddress == true && update.Message != nil {
		checkout.Address = update.Message.Text
		NewOrder(update, bot)
		return
	}
}

func NewOrder(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatId := api.GetChatId(update, bot)
	clientId := clients.GetClient(chatId)
	orderId := orders.NewOrder(chatId, clientId)

	if orderId > 0 {
		messageSuccess := "Заказ оформлен. Номер заказа: " + strconv.Itoa(orderId)
		message.Send(update, bot, messageSuccess)
	}

}
