package clients

import (
	"gitlab.com/chertokdmitry/bot-pandora-11/src/checkout"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/datasources/products_db"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
	"time"
)

const (
	queryGetClientByChatId = "SELECT id from clients WHERE chat_id = ?"
	queryInsertNewClient   = "INSERT INTO clients(chat_id, name, phone, address, created_at) VALUES (?, ?, ?, ?, ?)"
)

func GetClient(chatId int) int {
	db := products_db.GetDB()
	defer db.Close()

	Client := Client{}

	err := db.QueryRow(queryGetClientByChatId, chatId).Scan(&Client.Id)
	if err != nil {
		logger.Error("error when trying to get Client by chatId", err)
	}

	if Client.Id > 0 {
		return Client.Id
	} else {
		return NewClient(chatId)
	}
}

func NewClient(chatId int) int {
	db := products_db.GetDB()
	defer db.Close()

	res, err := db.Exec(queryInsertNewClient, chatId, checkout.Name, checkout.Phone, checkout.Address, time.Now())

	if err != nil {
		logger.Error("error when trying to insert new client", err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		logger.Error("error when trying to get new client id", err)
	}

	return int(lastId)
}
