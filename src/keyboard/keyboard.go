package keyboard

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/domain/categories"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
)

type Request struct {
	Action string `json:"action"`
	Id     string `json:"id"`
}

var MainPage = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Каталог"),
		tgbotapi.NewKeyboardButton("Корзина"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Заказы"),
		tgbotapi.NewKeyboardButton("Оформить заказ"),
	),
)

func GetCatalog(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	row1 := []tgbotapi.InlineKeyboardButton{}
	row2 := []tgbotapi.InlineKeyboardButton{}
	row3 := []tgbotapi.InlineKeyboardButton{}

	for _, category := range categories.CustomCategories {
		request := &Request{
			Action: "category",
			Id:     category.Id,
		}
		var msg, _ = json.Marshal(request)
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(category.Name, string(msg)))

		if category.Id == "22" {
			row1 = append(row1, row[0])
		}

		if category.Id == "20" {
			row2 = append(row2, row[0])
		}

		if category.Id == "21" {
			row3 = append(row3, row[0])
		}

	}

	msgInline := tgbotapi.NewMessage(int64(api.GetChatId(update, bot)), "Каталог")
	msgInline.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3)

	if _, err := bot.Send(msgInline); err != nil {
		logger.Error("error when trying to send catalog", err)
	}
}

func BackCatalog() tgbotapi.InlineKeyboardMarkup {
	requestData := &Request{
		Action: "back_catalog",
		Id:     "",
	}

	msgData, _ := json.Marshal(requestData)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вернуться", string(msgData)),
		),
	)
}

func AddToCart(productId string) tgbotapi.InlineKeyboardMarkup {
	requestData := &Request{
		Action: "add_product",
		Id:     productId,
	}

	msgData, _ := json.Marshal(requestData)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", string(msgData)),
		),
	)
}

func DeleteFromCart(productId string) tgbotapi.InlineKeyboardMarkup {
	requestData := &Request{
		Action: "delete_product",
		Id:     productId,
	}

	msgData, _ := json.Marshal(requestData)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить из корзины", string(msgData)),
		),
	)
}

//tgbotapi.NewInlineKeyboardRow(
//	tgbotapi.NewInlineKeyboardButtonURL("1.com","http://1.com"),
//	tgbotapi.NewInlineKeyboardButtonSwitch("2sw","open 2"),
//	tgbotapi.NewInlineKeyboardButtonData("3","3"),
//),
