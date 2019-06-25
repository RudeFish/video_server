package main

import (
	"net/http"
	"imooc/Go语言实战流媒体视频网站/video_server/api/session"
	"imooc/Go语言实战流媒体视频网站/video_server/api/defs"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

// 检验用户session合法性
func validateUserSession(r *http.Request) bool {
	// 获取id
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	// 判断id是否过期
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	// 存入session
	r.Header.Set(HEADER_FIELD_UNAME, uname)
	return true
}

// 检验用户合法性
func validateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}

	return true
}