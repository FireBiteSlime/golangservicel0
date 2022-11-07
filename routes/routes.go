package routes

import (
	"L0/cache"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDataByUid(cc *cache.Cache) gin.HandlerFunc {
	f := func(c *gin.Context) {
		uid := c.Param("uid")
		fmt.Println(c.Param("uid"))
		cacheOrder, status := cc.Get(uid)
		if status {
			c.JSON(http.StatusOK, cacheOrder)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Нет записей",
			})
		}

	}
	return gin.HandlerFunc(f)
}
