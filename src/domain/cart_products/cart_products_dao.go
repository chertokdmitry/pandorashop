package cart_products

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/datasources/products_db"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/message"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
	"time"
)

const (
	queryInsertCartProducts    = "INSERT INTO cart_products(cart_id, product_id, price, amount, created_at) VALUES (?, ?, ?, ?, ?)"
	queryDeleteCartProducts    = "DELETE FROM cart_products WHERE cart_id = ? AND product_id = ?"
	queryDeleteAllCartProducts = "DELETE FROM cart_products WHERE cart_id = ?"
)

type CartProduct struct {
	ProductId int
	Price     int
	Amount    int
}

func Insert(update tgbotapi.Update, bot *tgbotapi.BotAPI, cartId int, productId int, price int, amount int) {
	db := products_db.GetDB()
	defer db.Close()

	_, err := db.Exec(queryInsertCartProducts, cartId, productId, price, amount, time.Now())

	if err != nil {
		logger.Error("error when trying to insert new cart_products", err)
	}

	message.Send(update, bot, "Товар добавлен в корзину")
}

func Delete(cartId int, productId int) {
	db := products_db.GetDB()
	defer db.Close()

	_, err := db.Exec(queryDeleteCartProducts, cartId, productId)

	if err != nil {
		logger.Error("error when trying to delete cart_products", err)
	}
}

func DeleteAll(cartId int) {
	db := products_db.GetDB()
	defer db.Close()

	_, err := db.Exec(queryDeleteAllCartProducts, cartId)

	if err != nil {
		logger.Error("error when trying to delete cart_products", err)
	}
}
