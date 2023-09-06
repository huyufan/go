package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 * 1024
)

func main() {
	r := RegisterHandlers()

	http.ListenAndServe(":9999", r)
}

func RegisterHandlers() *httprouter.Router {
	route := httprouter.New()

	route.POST("/upload", uploadHandler)

	return route
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Printf("File is too big")
		return
	}
	files := r.MultipartForm.File["file"]

	num := len(files)

	fmt.Printf("总文件数：%d 个文件", num)

	//循环对每个文件进行处理
	for n, fheader := range files {
		//获取文件名
		filename := fheader.Filename

		//结束文件
		file, err := fheader.Open()

		fmt.Println(fheader.Size)
		if err != nil {
			fmt.Println(err)
		}

		//保存文件
		defer file.Close()
		f, err := os.Create("./video/" + filename)
		fmt.Println(f)
		defer f.Close()
		io.Copy(f, file)

		//获取文件状态信息
		fstat, _ := f.Stat()

		//打印接收信息
		fmt.Fprintf(w, "%s  NO.: %d  Size: %d KB  Name：%s\n", time.Now().Format("2006-01-02 15:04:05"), n, fstat.Size()/1024, filename)
		fmt.Printf("%s  NO.: %d  Size: %d KB  Name：%s\n", time.Now().Format("2006-01-02 15:04:05"), n, fstat.Size()/1024, filename)

	}

	// file, headers, err := r.FormFile("file[]")
	// if err != nil {
	// 	log.Printf("Error when try to get file: %v", err)
	// 	return
	// }
	// //获取上传文件的类型
	// if headers.Header.Get("Content-Type") != "image/png" {
	// 	log.Printf("只允许上传png图片")
	// 	return
	// }
	// data, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Printf("Read file error: %v", err)
	// 	return
	// }
	// fn := headers.Filename
	// err = os.WriteFile("./video/"+fn, data, 0666)
	// if err != nil {
	// 	log.Printf("Write file error: %v", err)
	// 	return
	// }
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
