package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"io"
	"encoding/json"
	"io/ioutil"
	"net/http/httputil"
	"net/url"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	// 从cookie中或取传过来的name
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	// 判断如果cookie中的username 或者sessionid有问题会跳转到登录页面
	if err1 != nil || err2 != nil{
		name := &HomePage{Name: "Silas"}
		t, err := template.ParseFiles("./template/home.html")
		if err != nil{
			log.Printf("Parsing template home.html error: %s\n", err)
			return
		}
		t.Execute(w, name)
		return
	}

	// 简单判断，如果用户名和sid都有调转到userhome页面
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "./userhome", http.StatusFound)
		return
	} else {
		io.WriteString(w, "err cookie!")
		return
	}
}

// userhome页面
func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	// 从cookie读取数据
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")

	// 判断如果cookie中的username 或者sessionid有问题会跳转到登录页面
	if err1 != nil || err2 != nil{
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// 从提交的表单读取数据
	// 这里不判断数据合法性，判断由前端调用api检测
	fname := r.FormValue("username")

	var p *UserPage
	if len(cname.Value)!= 0 {
		p = &UserPage{Name: cname.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}
	t, e := template.ParseFiles("./template/userhome.html")
	if e != nil{
		log.Printf("Parsing username.html error: %s\n", e)
		return
	}
	t.Execute(w, p)

}

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	// 方法应为post请求
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	// 解析body出错
	if err := json.Unmarshal(res, apibody); err != nil{
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}

	// 解析body
	request(apibody, w, r)
	defer r.Body.Close()
}

// 实现代理转发
func proxyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	u, _ := url.Parse("http://127.0.0.1:25602")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}