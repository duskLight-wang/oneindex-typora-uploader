package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func postFile(filename string, targetUrl string, fileName string) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		os.Exit(-1)
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		os.Exit(-1)
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		os.Exit(-1)
	}

	tr := &http.Transport{
		// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bodyWriter.Close()
	req, err := http.NewRequest("POST", targetUrl, bodyBuf)
	if err != nil {
		fmt.Println("发送报文失败", err)
		os.Exit(-1)
	}

	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/87.0")
	client := &http.Client{Transport: tr, Timeout: time.Second * time.Duration(60)}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(-1)
	}
	// fmt.Println(resp.Status)
	// fmt.Println(string(resp_body))
	// 获取返回的HTML内容
	str1 := string(resp_body)
	// 获取出现http的最后一次出现
	num1 := strings.LastIndex(str1, "http")
	// 获取文件名的最后一次出现
	num2 := strings.LastIndex(str1, fileName)
	fmt.Println(str1[num1 : num2+len(fileName)])
}

func main() {
	// 此处更换为你的oneindex域名
	tar_url := "https://example.com/images"
	var paths string
	var fileName string
	list := os.Args
	if len(list) == 2 {
		paths = list[1]
		// 获取文件名
		fileName = filepath.Base(paths)
	} else {
		return
	}
	postFile(paths, tar_url, fileName)
}
