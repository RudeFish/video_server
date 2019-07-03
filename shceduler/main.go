package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"imooc/Go语言实战流媒体视频网站/video_server/shceduler/taskrunner"
)

func RegisterHanders() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main()  {
	go taskrunner.Start()

	r := RegisterHanders()
	http.ListenAndServe(":25601", r)
}