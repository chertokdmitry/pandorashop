package routers

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/controllers"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/keyboard"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
	"strconv"
)

func RouteInline(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	data := keyboard.Request{}

	if err := json.Unmarshal([]byte(update.CallbackQuery.Data), &data); err != nil {
		logger.Error("error when trying to unmarshal callback data", err)
	}

	switch data.Action {
	case "add_product":
		productId, _ := strconv.Atoi(data.Id)
		controllers.AddToCart(update, bot, productId, api.GetChatId(update, bot))

	case "delete_product":
		productId, _ := strconv.Atoi(data.Id)
		controllers.DeleteFromCart(productId, api.GetChatId(update, bot))
		controllers.GetCart(update, bot)

	case "category":
		categoryId, _ := strconv.Atoi(data.Id)
		controllers.GetProductsByCategory(categoryId, update, bot)

	case "back_catalog":
		controllers.GetCatalog(update, bot)
	}
}
