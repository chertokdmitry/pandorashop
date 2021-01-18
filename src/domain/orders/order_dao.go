package orders

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/api"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/checkout"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/datasources/products_db"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/domain/cart_products"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/domain/carts"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/message"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
	"strconv"
	"time"
)

const (
	queryGetTotal          = "SELECT sum(price) FROM cart_products WHERE cart_products.cart_id = ?"
	queryInsertNewOrder    = "INSERT INTO orders(client_id, total, status_id, created_at) VALUES (?, ?, ?, ?)"
	queryViewOrderProducts = "SELECT product_id, amount FROM order_products WHERE order_id = ?"
	queryViewOrders        = "SELECT orders.id, total, status_id FROM orders, clients WHERE orders.client_id = clients.id AND clients.chat_id = ?"
)

func NewOrder(chatId int, clientId int) int {
	db := products_db.GetDB()
	defer db.Close()
	var total int

	cart := carts.GetCart(chatId)
	err := db.QueryRow(queryGetTotal, cart.Id).Scan(&total)

	if err != nil {
		logger.Error("error when trying to get total by cartId", err)
	}

	res, err := db.Exec(queryInsertNewOrder, clientId, total, 1, time.Now())

	if err != nil {
		logger.Error("error when trying to insert new order", err)
	}

	cart_products.DeleteAll(cart.Id)
	checkout.VarsFlash()
	lastId, err := res.LastInsertId()

	return int(lastId)
}

func ShowOrders(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatId := api.GetChatId(update, bot)

	db := products_db.GetDB()
	defer db.Close()

	data, err := db.Query(queryViewOrders, chatId)
	if err != nil {
		logger.Error("error when trying to get orders", err)
	}

	order := Order{}

	message.Send(update, bot, "Список заказов:")

	for data.Next() {
		err = data.Scan(&order.Id, &order.Total, &order.StatusId)
		if err != nil {
			logger.Error("error when trying to scan orders", err)
		}

		orderMessage := "Заказ N" + strconv.Itoa(order.Id) + ", на сумму: " + strconv.Itoa(order.Total) + "руб"
		message.Send(update, bot, orderMessage)
	}

}
