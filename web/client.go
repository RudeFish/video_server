package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"io"
	"bytes"
)

// 初始化一个client
var httpClient *http.Client

func init()  {
	httpClient = &http.Client{}
}


func request(b *ApiBody, w http.ResponseWriter, r *http.Request)  {
	var resp *http.Response
	var err error

	switch  b.Method{
	// 当方法位GET时
	case http.MethodGet:
		//Add request headers and body
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header

		resp, err = httpClient.Do(req)
		if err != nil{
			log.Printf(err.Error())
			return
		}
		normalResponse(w, resp)

	// processing post requests
	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewReader([]byte(b.ReqBody)))
		req.Header = r.Header

		resp, err = httpClient.Do(req)
		if err != nil{
			log.Printf(err.Error())
			return
		}
		normalResponse(w, resp)

	case http.MethodDelete:
		req, _ := http.NewRequest("Delete", b.Url, nil)
		req.Header = r.Header

		resp, err = httpClient.Do(req)
		if err != nil{
			log.Printf(err.Error())
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request!")
		return
	}
}


// Normal request processing
func normalResponse(w http.ResponseWriter, r *http.Response)  {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil{
		re, _ := json.Marshal(ErrorInternalFaults)
		io.WriteString(w, string(re))
		w.WriteHeader(500)
		return
	}

	io.WriteString(w, string(res))
	w.WriteHeader(r.StatusCode)
	return
}