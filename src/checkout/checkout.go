package checkout

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/message"
)

var IsName, IsPhone, IsAddress bool
var Name, Phone, Address string

func GetName(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	IsName = true
	IsPhone = false
	IsAddress = false

	message.Send(update, bot, "Введите имя:")
}

func GetPhone(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	IsName = false
	IsPhone = true
	IsAddress = false

	message.Send(update, bot, "Контактный телефон:")
}

func GetAddress(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	IsName = false
	IsPhone = false
	IsAddress = true

	message.Send(update, bot, "Адрес доставки:")
}

func VarsFlash() {
	IsName = false
	IsPhone = false
	IsAddress = false
}
