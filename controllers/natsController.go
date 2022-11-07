package controllers

import (
	"L0/cache"
	"L0/db"
	"L0/models"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/nats-io/stan.go"
)

func Sub(cache *cache.Cache, database *sql.DB) func(m *stan.Msg) {
	f := func(m *stan.Msg) {
		json := string(m.Data)
		order, err := UnmarshalJson(json)
		if err != nil {
			fmt.Println(err)
		} else {
			CheckCache(cache, order, database)
		}
	}
	return f
}

func CheckCache(cache *cache.Cache, order models.Order, database *sql.DB) {
	if _, state := cache.Get(order.Order_uid); !state {
		cache.Set(order.Order_uid, order)
		db.SetOrders(order, database)
	} else {
		if reflect.DeepEqual(order, db.GetOrder(database, order.Order_uid)) {
			fmt.Println("duplicate key: ", order.Order_uid)
		} else {
			db.UpdateOrder(order, database)
		}
	}
}
