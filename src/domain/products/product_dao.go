package products

import (
	"gitlab.com/chertokdmitry/bot-pandora-11/src/datasources/products_db"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
)

const (
	queryGetProductsByCategory = `SELECT id, name, url, stock, category_id, price FROM products WHERE stock > 0 AND category_id = ?`
	queryGetProductPrice       = `SELECT price FROM products WHERE id = ?`
	queryGetProductsByCartId   = `SELECT t1.id, t1.name, t1.url, t2.amount, t1.category_id, t1.price FROM cart_products t2 LEFT JOIN  products t1 ON t1.id = t2.product_id AND t1.category_id != 16 AND t2.cart_id = ?`
)

func GetPrice(productId int) int {
	db := products_db.GetDB()
	defer db.Close()

	var price int

	err := db.QueryRow(queryGetProductPrice, productId).Scan(&price)
	if err != nil {
		logger.Error("error when trying to get price by productId", err)
	}

	return price
}

func GetByCategory(catId int) []Product {
	db := products_db.GetDB()
	defer db.Close()

	data, err := db.Query(queryGetProductsByCategory, catId)
	if err != nil {
		logger.Error("error when trying to get products by category", err)
	}

	product := Product{}
	var products []Product

	for data.Next() {
		err = data.Scan(&product.Id, &product.Name, &product.Url, &product.Stock, &product.CategoryId, &product.Price)
		if err != nil {
			logger.Error("error when trying to scan products", err)
		}

		products = append(products, product)
	}

	return products
}

func GetByCart(cartId int) []Product {
	db := products_db.GetDB()
	defer db.Close()

	data, err := db.Query(queryGetProductsByCartId, cartId)
	if err != nil {
		logger.Error("error when trying to get products by cart id", err)
	}

	product := Product{}
	var products []Product

	for data.Next() {
		err = data.Scan(&product.Id, &product.Name, &product.Url, &product.Stock, &product.CategoryId, &product.Price)
		if err != nil {
			logger.Error("error when trying to scan products", err)
		}

		products = append(products, product)
	}

	return products
}
