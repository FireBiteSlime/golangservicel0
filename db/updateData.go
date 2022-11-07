package db

import (
	"L0/models"
	"database/sql"
)

func UpdateDelivery(delivery models.Delivery, database *sql.DB, order_id string) {
	update := `update "deliveries" set "Name"=$1, "Phone"=$2, "Zip"=$3, "City"=$4, "Address"=$5, "Region"=$6, "Email"=$7 where "Order_id"=$8`
	_, err := database.Exec(update, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email, order_id)
	CheckError(err)
}

func UpdatePayment(payment models.Payment, database *sql.DB) {
	update := `update "payments" set "Request_id"=$2, "Currency"=$3, "Provider"=$4, "Amount"=$5, "Payment_dt"=$6, "Bank"=$7, "Delivery_cost"=$8, "Goods_total"=$9, "Custom_fee"=$10 where "Transaction"=$1`
	_, err := database.Exec(update, payment.Transaction, payment.Request_id, payment.Currency, payment.Provider, payment.Amount, payment.Payment_dt, payment.Bank, payment.Delivery_cost, payment.Goods_total, payment.Custom_fee)
	CheckError(err)
}

func UpdateItems(items []models.Item, database *sql.DB) {
	update := `update "items" set "Track_number"=$2, "Price"=$3, "Rid"=$4, "Name"=$5, "Sale"=$6, "Size"=$7, "Total_price"=$8, "Nm_id"=$9, "Brand"=$10, "Status"=$11 where "Chrt_id"=$1 `
	for _, i := range items {
		_, err := database.Exec(update, i.Chrt_id, i.Track_number, i.Price, i.Rid, i.Name, i.Sale, i.Size, i.Total_price, i.Nm_id, i.Brand, i.Status)
		CheckError(err)
	}
}

func UpdateItem(item models.Item, database *sql.DB) {
	update := `update "items" set "Track_number"=$2, "Price"=$3, "Rid"=$4, "Name"=$5, "Sale"=$6, "Size"=$7, "Total_price"=$8, "Nm_id"=$9, "Brand"=$10, "Status"=$11 where "Chrt_id"=$1 `
	_, err := database.Exec(update, item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.Total_price, item.Nm_id, item.Brand, item.Status)
	CheckError(err)
}

func UpdateOrder(order models.Order, database *sql.DB) {
	update := `update "orders" set "Track_number"=$2, "Entry"=$3, "Locale"=$4, "Internal_signature"=$5, "Customer_id"=$6, "Delivery_service"=$7, "Shardkey"=$8, "Sm_id"=$9, "Date_created"=$10, "Oof_shard"=$11 where "Order_uid"=$1`
	_, err := database.Exec(update, order.Order_uid, order.Track_number, order.Entry, order.Locale, order.Internal_signature, order.Customer_id, order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard)
	CheckError(err)
	if !IsDelivery(database, order.Order_uid) {
		UpdateDelivery(order.Delivery, database, order.Order_uid)
	} else {
		SetDelivery(order.Delivery, database, order.Order_uid)
	}
	if !IsPayment(database, order.Order_uid) {
		UpdatePayment(order.Payment, database)
	} else {
		SetPayments(order.Payment, database)
	}
	for _, v := range order.Items {
		if !IsItem(database, v.Chrt_id) {
			UpdateItem(v, database)
		} else {
			SetItem(v, database)
		}
	}

}
