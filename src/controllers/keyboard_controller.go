package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/checkout"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/domain/orders"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/keyboard"
)

func GetCatalog(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	keyboard.GetCatalog(update, bot)
}

func GetCart(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	GetProductsByCart(update, bot)
}

func ShowOrders(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	orders.ShowOrders(update, bot)
}

func Checkout(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	checkout.VarsFlash()
	checkout.GetName(update, bot)
}
