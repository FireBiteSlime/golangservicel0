package db

import (
	"L0/models"
	"database/sql"
)

func SetDelivery(delivery models.Delivery, db *sql.DB, order_id string) {

	insert := `insert into "deliveries"("Name", "Phone","Zip", "City", "Address", "Region", "Email", "Order_id") values($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(insert, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email, order_id)
	CheckError(err)
}

func SetPayments(payment models.Payment, db *sql.DB) {
	insert := `insert into "payments"("Transaction", "Request_id","Currency", "Provider", "Amount", "Payment_dt", "Bank", "Delivery_cost","Goods_total","Custom_fee" ) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := db.Exec(insert, payment.Transaction, payment.Request_id, payment.Currency, payment.Provider, payment.Amount, payment.Payment_dt, payment.Bank, payment.Delivery_cost, payment.Goods_total, payment.Custom_fee)
	CheckError(err)
}

func SetItems(items []models.Item, db *sql.DB) {

	insert := `insert into "items"("Chrt_id", "Track_number","Price", "Rid", "Name", "Sale", "Size", "Total_price","Nm_id","Brand","Status" ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	for _, i := range items {

		_, err := db.Exec(insert, i.Chrt_id, i.Track_number, i.Price, i.Rid, i.Name, i.Sale, i.Size, i.Total_price, i.Nm_id, i.Brand, i.Status)
		CheckError(err)
	}
}

func SetItem(item models.Item, db *sql.DB) {
	insert := `insert into "items"("Chrt_id", "Track_number","Price", "Rid", "Name", "Sale", "Size", "Total_price","Nm_id","Brand","Status" ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(insert, item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.Total_price, item.Nm_id, item.Brand, item.Status)
	CheckError(err)
}

func SetOrders(order models.Order, db *sql.DB) {
	insert := `insert into "orders"("Order_uid", "Track_number","Entry", "Locale", "Internal_signature", "Customer_id", "Delivery_service", "Shardkey", "Sm_id", "Date_created", "Oof_shard") values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(insert, order.Order_uid, order.Track_number, order.Entry, order.Locale, order.Internal_signature, order.Customer_id, order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard)
	CheckError(err)
	SetDelivery(order.Delivery, db, order.Order_uid)
	SetPayments(order.Payment, db)
	SetItems(order.Items, db)
}
