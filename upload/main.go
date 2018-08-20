package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"io"
	"io/ioutil"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!\n")
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// 根据字段名获取表单文件
	formFile, header, err := r.FormFile("uploadfile")
	if err != nil {
		log.Printf("Get form file failed: %s\n", err)
		fmt.Fprintf(w, "Hello there!\n")
		return
	}
	defer formFile.Close()

	// 创建保存文件
	//destFile, err := os.Create("." + r.URL.Path + "/" + header.Filename)
	destFile, err := os.Create("." + "/" + header.Filename)
	if err != nil {
		log.Printf("Create failed: %s\n", err)
		return
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("Write file failed: %s\n", err)
		return
	}
}

func uploadBlob(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	ioutil.WriteFile("3.png", data, os.ModePerm)
}

func main() {
	tmp:= []byte("我是文本文件内容");
	ioutil.WriteFile("1.txt", tmp, os.ModePerm)

	http.HandleFunc("/", myHandler)
	http.HandleFunc("/upload", uploadHandle)
	http.HandleFunc("/upload_blob", uploadBlob)
	log.Fatal(http.ListenAndServe(":9998", nil))
	fmt.Println("Hello, 世界")
}