package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	meta "github.com/stsfang/go-netdisk/meta-dao"
	"github.com/stsfang/go-netdisk/util"
)

//UploadHandler 处理简单文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hello world"))

	//如果是Get方法，返回上传页面
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "Internal server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, filehead, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("fail to get data, err %s\n", err.Error())
			return
		}
		defer file.Close()

		//创建文件元数据，读取文件存放到本地文件系统
		fileMeta := meta.FileMeta{
			FileName: filehead.Filename,
			Location: "./tmp/" + filehead.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		//创建文件

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("创建文件失败:%s\n", err.Error())
			return
		}

		defer newFile.Close()

		//将文件拷贝到本地
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("拷贝文件失败 %s\n", err.Error())
			return
		}
		//文件游标重定位到文件起始，计算文件摘要
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)

		//TODO: 上传到ceph集群
		//文件游标重定位到文件起始，读取文件
		//newFile.Seek(0, 0)
		//data, err = ioutil.ReadAll(newFile)
		// if err != nil {
		// 	fmt.Printf("读取文件失败 %s\n", err.Error())
		// 	return
		// }

		//TODO: 上传到aliyun oss

		//TODO: 更新数据库表
		fmt.Println("add...")
		meta.AddFileMeta(fileMeta)
		fmt.Printf("%v\n", fileMeta)

		//返回响应
		http.Redirect(w, r, "/static/view/home.html", http.StatusFound)
	}
}
