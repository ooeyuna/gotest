package http

import (
	"net/http"
)

func Install(){
	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		res.Write([]byte("hello world"))
	})
	http.HandleFunc("/test",func(res http.ResponseWriter,req *http.Request){
		res.Write([]byte("req a:"+ req.FormValue("a")))
	})
	http.ListenAndServe("localhost:8091",nil);
}