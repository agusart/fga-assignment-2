package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func StartRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})

	}

	return router
}
