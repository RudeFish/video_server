package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


// 添加middleware
type middleWareHandler struct {
	r *httprouter.Router
}


func NewMinddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

// 将middleWareHandler实现handler方法
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	// check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

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
	mh := NewMinddleWareHandler(r)
	http.ListenAndServe(":25600", mh)
}
