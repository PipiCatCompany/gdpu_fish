package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	v1 "go-xianyu/api/v1"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

func SyncMessageToCpp(v *viper.Viper, message v1.CreateMessageRequest) error {

	// cppIp := v.GetString("cpp.host")
	// cppPort := v.GetString("cpp.port")

	url := "http://192.168.1.110:10087/message"
	// url := "http://" + cppIp + ":" + cppPort + "/message"
	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshaling message to JSON: %w", err)
	}

	// 创建一个新的请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		// 读取并打印错误响应内容
		_, err := io.ReadAll(resp.Body) // 使用 io.ReadAll
		return err
	}

	return nil
}
