package app

import (
	"gitlab.com/chertokdmitry/bot-pandora-11/src/api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/controllers"
)

var Name string

func StartApplication() {
	bot := api.NewBot()

	for update := range api.Update(bot) {
		controllers.Handler(update, bot)

		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}
		Map(update, bot)
	}
}
