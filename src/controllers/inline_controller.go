package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/domain/cart_products"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/domain/carts"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/domain/products"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/keyboard"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
	"strconv"
)

func GetProductsByCategory(key int, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatId := api.GetChatId(update, bot)
	data := products.GetByCategory(key)
	SendProductsByCategory(update, bot, data, chatId)
}

func GetProductsByCart(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatId := api.GetChatId(update, bot)
	cart := carts.GetCart(chatId)
	data := products.GetByCart(cart.Id)
	SendProductsByCart(update, bot, data, chatId)
}

func SendProductsByCategory(update tgbotapi.Update, bot *tgbotapi.BotAPI, data []products.Product, chatId int) {
	for _, product := range data {
		file := "http://charmclub.ru/wp-content/uploads/" + product.Url
		caption := product.Name + "\n" + "Цена: " + strconv.Itoa(product.Price) + " руб"

		msgFile := tgbotapi.NewPhotoUpload(int64(api.GetChatId(update, bot)), nil)
		msgFile.FileID = file
		msgFile.UseExisting = true
		msgFile.Caption = caption
		msgFile.ReplyMarkup = keyboard.AddToCart(strconv.Itoa(product.Id))

		if _, err := bot.Send(msgFile); err != nil {
			logger.Error("error when trying to send to bot product", err)
		}
	}
}

func SendProductsByCart(update tgbotapi.Update, bot *tgbotapi.BotAPI, data []products.Product, chatId int) {
	for _, product := range data {
		file := "http://charmclub.ru/wp-content/uploads/" + product.Url
		caption := product.Name + "\n" + "Цена: " + strconv.Itoa(product.Price) + " руб" + "\n" + "Количество: " + strconv.Itoa(product.Stock)

		msgFile := tgbotapi.NewPhotoUpload(int64(chatId), nil)
		msgFile.FileID = file
		msgFile.UseExisting = true
		msgFile.Caption = caption
		msgFile.ReplyMarkup = keyboard.DeleteFromCart(strconv.Itoa(product.Id))

		if _, err := bot.Send(msgFile); err != nil {
			logger.Error("error when trying to send to bot product", err)
		}
	}

	msgInline := tgbotapi.NewMessage(int64(chatId), "Вернуться")
	msgInline.ReplyMarkup = keyboard.BackCatalog()
	if _, err := bot.Send(msgInline); err != nil {
		logger.Error("error when trying to send to bot button back", err)
	}
}

func AddToCart(update tgbotapi.Update, bot *tgbotapi.BotAPI, productId int, chatId int) {
	cart := carts.GetCart(chatId)
	price := products.GetPrice(productId)
	cart_products.Insert(update, bot, cart.Id, productId, price, 1)
}

func DeleteFromCart(productId int, chatId int) {
	cart := carts.GetCart(chatId)
	cart_products.Delete(cart.Id, productId)
}
