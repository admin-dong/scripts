package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//截图之后 然后就是上传了 校验

func T_token() (string, error) {
	url := "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal"
	payload := map[string]string{
		"app_id":     "cli_a5fc3ff4e07e500c",
		"app_secret": "wKSgxDQDZed7lvE9WayLzdphDO5q5GKp",
	}

	// jsonPayload, err := json.Marshal(payload)
	// if err != nil {
	// 	fmt.Println("序列化失败")
	// 	return "", err
	// }
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("序列化失败")
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("创建请求失败", err)
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("发送请求失败", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取响应失败", err)
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("解析失败", err)
		return "", err
	}
	token := data["tenant_access_token"].(string)
	fmt.Println("获取t_token成功")
	fmt.Println(token)
	return token, err

}
