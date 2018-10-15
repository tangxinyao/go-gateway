package domain

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func StatusOK(context echo.Context, data interface{}) error {
	response := new(Response)
	response.Status = http.StatusOK
	response.Data = data
	return context.JSON(http.StatusOK, response)
}

func StatusInternalServerError(context echo.Context, err error) error {
	response := new(Response)
	response.Status = http.StatusInternalServerError
	response.Message = err.Error()
	return context.JSON(http.StatusOK, response)
}

func StatusUnauthorized(context echo.Context, err error) error {
	response := new(Response)
	response.Status = http.StatusUnauthorized
	response.Message = "用户名或密码错误"
	return context.JSON(http.StatusOK, response)
}
