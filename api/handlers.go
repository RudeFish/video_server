package main

import (
	"net/http"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"io"
	"imooc/Go语言实战流媒体视频网站/video_server/api/defs"
	"encoding/json"
	"imooc/Go语言实战流媒体视频网站/video_server/api/dbops"
	"imooc/Go语言实战流媒体视频网站/video_server/api/session"
)

// 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	// 读取body内容
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	// 错误部分处理
	// 反序列化json
	if err := json.Unmarshal(res, ubody); err != nil{
		sendErrorResponse(w, defs.ErrorRequestsBodyParseFaild)
		return
	}
	// db添加用户
	if err := dbops.AddUserCredential(ubody.UserName, ubody.Pwd); err != nil{
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	// 校验正确
	id := session.GenerateNewSessionId(ubody.UserName)
	su := defs.SignedUp{true, id}

	if resp, err := json.Marshal(su); err != nil{
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

// 用户登录,带参数
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)   {
	// 接收参数
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}