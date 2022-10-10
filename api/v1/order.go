package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewOrderRequqest struct {
}

type NewOrderItem struct {
}

func CreateNewOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

	}
}
