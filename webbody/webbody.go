package webbody

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"durian/helpers"
)

var ProxyUrl string

func WebSetProxy(proxuUrl string) {
	ProxyUrl = proxuUrl
}

func WebGet(method, uri string) (error, []byte) {
	// 设置代理地址
	proxyUrl, err := url.Parse(ProxyUrl)
	if err != nil {
		return err, nil
	}

	// 创建Transport对象，并设置代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	// 创建Client对象，并设置Transport
	client := &http.Client{
		Transport: transport,
		Timeout:   12 * time.Second,
	}

	// 构造HTTP请求
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return err, nil
	}

	// 发送HTTP请求，获取响应
	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}

	// 读取响应内容
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	return nil, body
}

func WriteFile(method, targetURL, savePath string) (error, string, int64) {
	sp := strings.Split(targetURL, "/")
	fileName := sp[len(sp)-1]
	if strings.Contains(fileName, "?") {
		spa := strings.Split(fileName, "?")
		fileName = spa[0]
	}
	if filepath.Ext(fileName) == "" {
		fileName = fileName + ".jpg"
	}
	dira := filepath.Join(savePath, fileName)
	if helpers.FileExists(dira) {
		return fmt.Errorf("File already exists."), "", 0
	}
	// 设置代理地址
	proxyUrl, err := url.Parse(ProxyUrl)
	if err != nil {
		return err, "", 0
	}

	// 创建Transport对象，并设置代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	// 创建Client对象，并设置Transport
	client := &http.Client{
		Transport: transport,
	}

	// 构造HTTP请求
	req, err := http.NewRequest(method, targetURL, nil)
	if err != nil {
		return err, "", 0
	}

	// 发送HTTP请求，获取响应
	resp, err := client.Do(req)
	if err != nil {
		return err, "", 0
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Server return error code:%v\n", resp.Status), "", 0
	}

	// sp := strings.Split(targetURL, "/")
	// fileName := sp[len(sp)-1]
	// dira := filepath.Join(savePath, fileName)
	// fmt.Println("Final imageFile:", dira)
	file, err := os.OpenFile(dira, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err, "", 0
	}
	defer file.Close()

	wbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		return err, "", 0
	}
	return nil, dira, wbytes
}
