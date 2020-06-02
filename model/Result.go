package model

import "net/http"

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

type PageParam struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
	Total    int `json:"total"`
}

func FailResult(message string, err error) Result {
	return Result{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    err.Error(),
	}
}

func SuccessResult(data interface{})  Result{
	return Result{
		Code: http.StatusOK,
		Message: "请求成功",
		Data: data,
	}
}
