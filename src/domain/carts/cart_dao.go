package carts

import (
	"gitlab.com/chertokdmitry/bot-pandora-11/src/datasources/products_db"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
	"time"
)

const (
	queryGetCartByChatId  = "SELECT id, chat_id, cart_actf from carts WHERE chat_id = ?"
	queryInsertNewCart    = "INSERT INTO carts(chat_id, cart_actf, created_at) VALUES (?, 1, ?)"
	queryViewCartProducts = "SELECT product_id, amount FROM cart_products WHERE cart_id = ?"
)

func GetCart(chatId int) Cart {
	db := products_db.GetDB()
	defer db.Close()

	cart := Cart{}

	err := db.QueryRow(queryGetCartByChatId, chatId).Scan(&cart.Id, &cart.ChatId, &cart.CartActf)
	if err != nil {
		logger.Error("error when trying to get cart by chatId", err)
	}

	if cart.Id > 0 {
		return cart
	} else {
		return NewCart(chatId)
	}
}

func NewCart(chatId int) Cart {
	db := products_db.GetDB()
	defer db.Close()

	res, err := db.Exec(queryInsertNewCart, chatId, time.Now())

	if err != nil {
		logger.Error("error when trying to insert new card", err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		logger.Error("error when trying to get new cart id", err)
	}

	cart := Cart{}

	cart.Id = int(lastId)
	cart.ChatId = chatId
	cart.CartActf = 1

	return cart
}
