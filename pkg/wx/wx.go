package wx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type GetOpenIdByCodeResponse struct {
	Session_key string
	Unionid     string
	Errmsg      string
	Openid      string
	Errcode     int32
}

func GetOpenIdByCode(code string) GetOpenIdByCodeResponse {
	appID := "wx09070093b5f169f0"                // 替换为你的 APPID
	secret := "ccdcae93597d9b584aafb2b1f1f3fb06" // 替换为你的 SECRET
	jsCode := code                               // 替换为你的 JSCODE

	// 构建请求 URL
	apiURL := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", url.QueryEscape(appID), url.QueryEscape(secret), url.QueryEscape(jsCode))

	// 创建 HTTP GET 请求
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v", err)
		return GetOpenIdByCodeResponse{}
	}

	var response GetOpenIdByCodeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error parsing JSON response: %v", err)
		return GetOpenIdByCodeResponse{}
	}

	return response
}
