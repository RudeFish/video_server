package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"time"
	"log"
	"io/ioutil"
	"io"
	"html/template"
)

// 测试视频上传
func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	t, err := template.ParseFiles("./videos/upload.html")
	if err != nil{
		log.Printf("Open html file error: %v\n", err)
		sendErrorResponse(w, http.StatusNotFound, "Page not found!")
		return
	}
	t.Execute(w, nil)
}


// 获取文件
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	// 获取文件id和路径
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil{
		log.Printf("Error when try to open file %v\n", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	// 加入header视频文件强制提醒
	w.Header().Set("Content-type", "video/mp4")
	// 传输二进制流 播放视频
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}


func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	// 限定上传文件的大小
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAN_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAN_SIZE); err != nil{
		log.Printf("Error when try upload file: %v\n", err)
		sendErrorResponse(w, http.StatusBadRequest, "File is too big!")
		return
	}

	// 读取表单
	file, _, err := r.FormFile("file") // key位html页面name值
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error!")
		return
	}

	// 读取文件
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file Error!")
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error!")
		return
	}

	// 文件名
	fn := p.ByName("vid-id")
	// 写入服务端
	err = ioutil.WriteFile(VIDEO_DIR + fn, data, 0666)
	if err != nil{
		log.Printf("Write file error: %v \n", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error!")
		return
	}

	// 上传成功
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload successfully!")
}