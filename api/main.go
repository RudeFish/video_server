package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHanders() *httprouter.Router {
	router := httprouter.New()
	// 创建用户
	router.POST("/user", CreateUser)
	// 用户登录,带参数
	router.POST("/user/:user_name", Login)


	return router
}


func main()  {
	r := RegisterHanders()
	http.ListenAndServe(":25600", r)
}
