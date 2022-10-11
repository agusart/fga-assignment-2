package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"tugas2/api/middleware"
	v1 "tugas2/api/v1"
)

func StartRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})

		api.GET("/orders", v1.GetOrders(db))
		api.Use(middleware.RequestMustBeJSON()).POST("/orders", v1.CreateNewOrder(db))

	}

	return router
}
