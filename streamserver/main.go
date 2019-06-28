package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type midelWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}


func NewMiddleWareHandle(r *httprouter.Router, cc int) http.Handler {
	m := midelWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m  midelWareHandler)ServeHTTP (w http.ResponseWriter, r *http.Request)  {
	// 判断如果超过流控值
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too Many Requests")
		return
	}

	m.r.ServeHTTP(w, r)
	// 释放token
	defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)

	// 测试
	router.GET("/testpage", testPageHandler)

	return router
}

func main()  {
	r := RegisterHandlers()
	m := NewMiddleWareHandle(r, 2)
	http.ListenAndServe(":25600", m)
}
