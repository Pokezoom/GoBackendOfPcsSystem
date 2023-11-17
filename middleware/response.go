package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 发http请求方法，调模块方法用
func SendHTTPRequest(method, url string, body []byte) ([]byte, error) {
	// 创建请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// 设置请求头，这里可以根据需要进行修改或添加更多头部信息
	req.Header.Set("Content-Type", "application/json")

	// 设置客户端以及超时时间
	client := &http.Client{
		Timeout: 15 * time.Second, // 设置超时时间，这里设置为15秒
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 确保响应体被关闭

	// 读取响应
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//func main() {
//	// 示例：发送GET请求
//	url := "http://example.com"
//	response, err := sendHTTPRequest("GET", url, nil)
//	if err != nil {
//		panic(err)
//	}
//	println(string(response))
//
//	// 如果需要发送POST请求，可以传递带有数据的body
//	// postBody := []byte(`{"key":"value"}`)
//	// response, err = sendHTTPRequest("POST", url, postBody)
//	// ...处理响应和错误
//}
