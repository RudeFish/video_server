package main

import (
	"net/http"
	"io"
)

// 出错时返回

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string)  {
	// 错误码写到handler，错误信息打印
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}