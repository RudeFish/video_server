package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"imooc/Go语言实战流媒体视频网站/video_server/api/session"
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
	// 获取当前用户信息
	router.GET("/user/:username", GetUserInfo)
	// 添加视频，body带上用户id和视频名字,注意authorid是int
	router.POST("/user/:user_name/videos", AddNewVideo)
/*
	router.GET("/user/:username/videos", ListAllVideos)

	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)

	router.POST("/videos/:vid-id/comments", PostComment)

	router.GET("/videos/:vid-id/comments", ShowComments)
*/
	return router
}

// 将数据库中所有session加载到程序的map中
func Prepare()  {
	session.LoadSessionFromDB()
}


func main()  {
	Prepare()
	r := RegisterHanders()
	mh := NewMinddleWareHandler(r)
	http.ListenAndServe(":25600", mh)
}
