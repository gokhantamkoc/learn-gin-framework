package main

import (
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateSuccessResponse(data interface{}) Response {
	response := Response{
		Status:  http.StatusOK,
		Success: true,
		Message: "",
		Data:    data,
	}
	return response
}

func CreateBadRequestResponse(err error) Response {
	response := Response{
		Status:  http.StatusBadRequest,
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}
	return response
}
