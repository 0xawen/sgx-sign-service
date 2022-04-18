package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

var (
	URL  = "http://127.0.0.1:8080"
	addr = "V6Avp9KqLfwGUaRFPV27a8VZhwxiYofoU"
	msg  = "123456"
)

// 集成测试
func main() {
	//fmt.Println("==========  测试: /ping ========== ")
	//GET(URL + "/ping")
	//
	//fmt.Println("==========  测试: /create ========== ")
	//fmt.Println(GET(URL + "/create"))

	fmt.Println("==========  测试: /sign ========== ")
	args := map[string]interface{}{
		"address": addr,
		"msg":     msg,
	}
	result := POST(URL+"/sign", args)
	sign, ok := result["data"]
	if !ok {
		fmt.Println("key data not exist")
	}
	fmt.Println(sign)

	fmt.Println("========== 测试: /verify ==========")
	args1 := map[string]interface{}{
		"address": addr,
		"sign":    sign,
		"msg":     msg,
	}
	result1 := POST(URL+"/verify", args1)
	fmt.Println(result1)
}

func GET(url string) map[string]string {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("GET")

	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return nil
	}
	b := resp.Body()
	fmt.Println("result:\r\n", string(b))

	result := map[string]string{}
	_ = json.Unmarshal(b, &result)
	return result
}

func POST(url string, args map[string]interface{}) map[string]interface{} {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)

	requestBody, err := json.Marshal(&args)
	if err != nil {
		fmt.Println(" paramer error")
		return nil
	}
	req.SetBody(requestBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return nil
	}
	b := resp.Body()
	fmt.Println("result:\r\n", string(b))

	result := map[string]interface{}{}
	_ = json.Unmarshal(b, &result)
	return result
}
