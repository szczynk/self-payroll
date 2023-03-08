package helper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	successJson struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	errorJson struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}

	ErrorWithCode struct {
		Msg  string `json:"message"`
		Code int    `json:"status"`
	}
)

func ResponseSuccessJson(c echo.Context, message string, data interface{}) error {

	if message == "" {
		message = "success"
	}

	res := successJson{
		Message: message,
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, res)
}

func ResponseValidationErrorJson(c echo.Context, message string, detail interface{}) error {
	res := errorJson{
		Message: message,
		Success: false,
		Error:   detail,
	}

	return c.JSON(http.StatusBadRequest, res)
}

func ResponseErrorJson(c echo.Context, code int, err error) error {
	res := errorJson{
		Error: err.Error(),
	}

	return c.JSON(code, res)
}
