package v1

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
	"tugas2/api/handler"
	"tugas2/model"
)

type NewOrderRequest struct {
	OrderedAt    string         `json:"orderedAt" valid:"required~orderedAt is required,rfc3339~unknown time format"`
	CustomerName string         `json:"customerName" valid:"required~customerName is required,alpha~name should not contain numeric and symbol"`
	Items        []NewOrderItem `json:"items"`
}

type NewOrderItem struct {
	ItemCode    string `json:"itemCode" valid:"required~itemCode is required,alphanum~itemCode must alphanumeric"`
	Description string `json:"description" valid:"optional"`
	Quantity    uint   `json:"quantity" valid:"required~quantity is required,range(1|999)~item quantity must between 1 to 999"`
}

func (orderRequest *NewOrderRequest) validate() error {
	_, err := govalidator.ValidateStruct(orderRequest)
	if err != nil {
		return err
	}

	if len(orderRequest.Items) < 1 {
		return errors.New("items is empty")
	}

	for _, item := range orderRequest.Items {
		_, err := govalidator.ValidateStruct(item)
		if err != nil {
			return err
		}
	}

	return nil
}

func (orderRequest *NewOrderRequest) NewOrderFromRequest() (*model.Order, error) {
	t, err := time.Parse(time.RFC3339, orderRequest.OrderedAt)
	if err != nil {
		return nil, err
	}
	order := &model.Order{
		OrderedAt:    t,
		CustomerName: orderRequest.CustomerName,
	}

	return order, nil
}

func CreateNewOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var orderRequest NewOrderRequest
		if err := c.ShouldBindJSON(&orderRequest); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "invalid json format",
				"code":    handler.BadRequestErrorCode,
			})

			log.Println(err.Error())
			return
		}

		if err := orderRequest.validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"code":    handler.ValidationError,
			})

			return
		}

		order, err := orderRequest.NewOrderFromRequest()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
				"code":    handler.ValidationError,
			})

			return
		}

		db.Create(&order)

		for _, item := range orderRequest.Items {
			orderItem := model.Item{
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
				OrderID:     order.ID,
			}

			order.Items = append(order.Items, orderItem)
		}

		db.Save(&order)

		c.JSON(http.StatusOK, gin.H{
			"success":  "true",
			"order_id": order.ID,
		})

		return
	}
}

type OrderResponse struct {
	OrderedAt    string              `json:"orderedAt"`
	CustomerName string              `json:"customerName"`
	OrderID      uint                `json:"orderID"`
	Items        []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	LineItemID  int    `json:"lineItemID,omitempty"`
}

func NewResponseFromModel(order model.Order) OrderResponse {
	var ordersResponse OrderResponse
	ordersResponse.Items = make([]OrderItemResponse, 0)

	ordersResponse.OrderID = order.ID
	ordersResponse.CustomerName = order.CustomerName
	ordersResponse.OrderedAt = order.OrderedAt.Format(time.RFC3339)

	for _, item := range order.Items {
		itemResponse := OrderItemResponse{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			LineItemID:  item.LineItemID,
		}
		ordersResponse.Items = append(ordersResponse.Items, itemResponse)
	}

	return ordersResponse
}

func GetOrders(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var orderResponse []OrderResponse

		var orders []model.Order
		db.Preload("Items").Find(&orders)

		for _, order := range orders {
			orderResponse = append(orderResponse, NewResponseFromModel(order))
		}

		c.JSON(http.StatusOK, gin.H{
			"orders": orderResponse,
			"count":  len(orderResponse),
		})

		return
	}
}
