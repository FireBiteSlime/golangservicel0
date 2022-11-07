package main

import (
	"L0/cache"
	"L0/controllers"
	"L0/db"
	"L0/models"
	"L0/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
)

func main() {

	preTime, _ := time.ParseDuration("20m")
	cache := cache.New() //кэш
	var orders []models.Order
	database := db.Connect() //база
	defer database.Close()

	for _, v := range orders {
		cache.Set(v.Order_uid, v)
	}

	sc, _ := stan.Connect("test-cluster", "test", stan.NatsURL("nats://host.docker.internal:4223"))
	defer sc.Close()
	sc.Subscribe("foo", controllers.Sub(cache, database), stan.StartAtTimeDelta(preTime))

	r := gin.Default() // сервер
	r.GET("/order/:uid", routes.GetDataByUid(cache))
	r.Run()
}
