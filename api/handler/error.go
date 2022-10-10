package handler

var (
	ValidationError = ValidationErrorResponse{
		Status:  "ER-422",
		Message: "validation-failed",
	}
)
