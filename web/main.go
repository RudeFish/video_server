package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	route := httprouter.New()
	route.GET("/", homeHandler)
	route.POST("/", homeHandler)

	route.GET("/userhome", userHomeHandler)
	route.POST("/userhome", userHomeHandler)

	route.POST("/api", apiHandler)

	route.POST("/upload/:vid-id", proxyHandler)

	// 将template文件夹下的内容绑定到static/下
	route.ServeFiles("/static/*filepath", http.Dir("./template"))

	return route
}

func main()  {
	r := RegisterHandlers()

	http.ListenAndServe(":25603", r)
}
