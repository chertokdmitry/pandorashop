package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
)

var botMessage string

func Get(updateText string) string {
	if botMessage != "" {
		answer := botMessage
		botMessage = ""

		return answer
	} else {
		return updateText
	}
}

func Set(message string) {
	botMessage = message
}

func Send(update tgbotapi.Update, bot *tgbotapi.BotAPI, message string) {
	msgMain := tgbotapi.NewMessage(int64(api.GetChatId(update, bot)), message)

	if _, err := bot.Send(msgMain); err != nil {
		logger.Error("error when trying to send main response message", err)
	}
}
