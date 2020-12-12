package http

import (
	"errors"
	thy "github.com/wozaifeiyang0/thylog"
	"net/http"
	"time"
)

/// 带重试功能的http请求
func Get(url string, retryNum , retryDelay , timeout int) (*http.Response, error) {

	errMsg := ""
	for index := 0; index < retryNum; index++ {

		req, _ := http.NewRequest("GET", url, nil)
		// 设置请求头，模拟浏览器访问
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

		// 发起请求
		resp, err := (&http.Client{Timeout: time.Second * time.Duration(timeout)}).Do(req)

		// 判断请求是否异常
		if err == nil && resp.StatusCode == 200 {
			return resp, err
		}else if err != nil {
			thy.Error.Printf("请求%c[31;4;31m %s %c[0m异常【状态码：%d】 :   %s    \n", 0x1B, "超时", 0x1B, 408, url)
			errMsg = err.Error()
		} else if resp.StatusCode == 404 {
			thy.Error.Printf("请求%c[31;4;31m %s %c[0m异常【状态码：%d】 :   %s    \n", 0x1B, "地址不存在", 0x1B, resp.StatusCode, url)
			errMsg = "地址异常：返回状态码404"
			return nil, errors.New("SendRetry err:" + errMsg)
		} else {
			thy.Error.Printf("请求%c[31;4;31m %s %c[0m异常【状态码：%d】 :   %s    \n", 0x1B, "服务访问", 0x1B, resp.StatusCode, url)
			errMsg = "地址异常：返回状态码不是200"
		}
		// 如果请求正常关闭
		if resp != nil{
			resp.Body.Close()
		}
		// 延迟几秒再次请求
		time.Sleep(time.Duration(retryDelay) * time.Second)
	}

	return nil, errors.New("SendRetry err:" + errMsg)
}
