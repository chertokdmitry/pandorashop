package order_products

import (
	"gitlab.com/chertokdmitry/bot-pandora-11/src/datasources/products_db"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
	"time"
)

const (
	queryInsertCartProducts = "INSERT INTO cart_products(cart_id, product_id, amount, created_at) VALUES (?, ?, ?, ?)"
	queryDeleteCartProducts = "DELETE FROM cart_products WHERE cart_id = ? AND product_id = ?"
)

type CartProduct struct {
	ProductId int
	Amount    int
}

func Insert(cartId int, productId int, amount int) {
	db := products_db.GetDB()
	defer db.Close()

	_, err := db.Exec(queryInsertCartProducts, cartId, productId, amount, time.Now())

	if err != nil {
		logger.Error("error when trying to insert new cart_products", err)
	}
}

func Delete(cartId int, productId int) {
	db := products_db.GetDB()
	defer db.Close()

	_, err := db.Exec(queryDeleteCartProducts, cartId, productId)

	if err != nil {
		logger.Error("error when trying to delete cart_products", err)
	}
}
