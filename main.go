package main

import (
	"net/http"
	"os"

	"github.com/stsfang/go-netdisk/handler"
)

func main() {
	//当前文件系统路径
	root, err := os.Getwd()
	if err != nil {
		panic("无法查看工作目录")
	}
	//处理简单文件上传
	http.HandleFunc("/file/upload", handler.UploadHandler)
	//静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(root+"/static"))))
	http.ListenAndServe(":8080", nil)
}
