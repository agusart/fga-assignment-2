package handler

import (
	"github.com/gin-gonic/gin"
)

type FieldsError struct {
	Name  string `json:"name"`
	Error string `json:"error"`
}

type ValidationErrorResponse struct {
	Status  string        `json:"code"`
	Message string        `json:"message"`
	Fields  []FieldsError `json:"fields"`
}

func SendResponse(c *gin.Context, httpStatus int, response any) {
	c.JSON(httpStatus, response)
}
