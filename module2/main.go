package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/healthz", myHandle)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println("系统运行异常", err)
	}

}

func myHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("开始处理请求...")
	makeResponseHeaders(w, r)
	//w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, "200")
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)

	log.Println("请求完成:请求IP:", r.RemoteAddr, "StatusCode:", http.StatusOK)
}

// 配置响应头
func makeResponseHeaders(w http.ResponseWriter, r *http.Request) {

	h := w.Header()

	//请求头内容设置到响应头
	var value strings.Builder
	for k, v := range r.Header {
		for i, s := range v {
			if i > 0 {
				value.WriteString(",")
			}
			value.WriteString(s)
		}
		h.Set(k, value.String())
		value.Reset()
	}

	//添加VERSION
	h.Set("VERSION", os.Getenv("VERSION"))

	log.Println("响应头配置完成！")
}
