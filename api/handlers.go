package main

import (
	"net/http"
	"io"
	"github.com/julienschmidt/httprouter"
)

// 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	io.WriteString(w, "test")
}

// 用户登录,带参数
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)   {
	// 接收参数
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}