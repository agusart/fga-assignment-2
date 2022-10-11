package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	OrderedAt    time.Time
	CustomerName string
	Items        []Item `gorm:"foreignKey:OrderID"`
}

type Item struct {
	gorm.Model
	ItemCode    string
	Description string
	Quantity    uint
	LineItemID  int
	OrderID     uint
}
