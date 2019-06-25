package main

import (
	"net/http"
	"imooc/Go语言实战流媒体视频网站/video_server/api/defs"
	"encoding/json"
	"io"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse)  {
	w.WriteHeader(errResp.HttpSc)

	resRtr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resRtr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int)  {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}