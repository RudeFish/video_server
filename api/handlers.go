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
	"log"
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
	log.Printf("Add username:%s, pwd:%s",ubody.UserName, ubody.Pwd)

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
	io.WriteString(w, uname + " Successful login!")
}


// 获取当前用户信息
func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params)   {
	// 判算用户合法性
	if err := !validateUser(w, r); err == true{
		log.Printf("Unauthorized user: %v\n", err)
		io.WriteString(w, "Unauthorized user!")
		return
	}
	uname := p.ByName("username")
	u, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	uid := &defs.UserInfo{Id: u.Id, Name: u.LoginName}
	if resp, err := json.Marshal(uid); err != nil{
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}else {
		sendNormalResponse(w, string(resp), 200)
	}
}



// ---------------------------------------------- video -------------------------------------------------------- //
func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !validateUser(w, r) {
		log.Printf("Unathorized user \n")
		return
	}
	// 读取body信息
	rep, _ := ioutil.ReadAll(r.Body)
	nvbody := &defs.VideoInfo{}

	if err := json.Unmarshal(rep, nvbody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestsBodyParseFaild)
		return
	}

	vi, err := dbops.AddNewVideo(nvbody.AuthorId, nvbody.Name)
	if err != nil{
		log.Printf("Error of AddNewVideo: %v\n", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if rep, err := json.Marshal(vi); err != nil{
		log.Printf("Error of AddNewVideo: %v\n", err)
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}else {
		sendNormalResponse(w, string(rep), 201)
	}
}