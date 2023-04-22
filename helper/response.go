package helper

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type JsonWrap struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

// OkJson 返回一个成功响应体
func OkJson(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, JsonWrap{0, data, "ok"})
}

// ErrorJson 返回一个错误信息响应体
func ErrorJson(c echo.Context, msg string) error {
	return c.JSON(http.StatusOK, JsonWrap{1, nil, msg})
}

// StatusErrorJson 返回一个自定义错误状态码的错误信息响应体
func StatusErrorJson(c echo.Context, msg string, status int) error {
	return c.JSON(http.StatusOK, JsonWrap{status, nil, msg})
}
