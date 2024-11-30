package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// 获取token之后的操作
func SendImage(tToken string) {
	url := "https://open.feishu.cn/open-apis/im/v1/images"

	// 打开图片文件
	file, err := os.Open("C:\\Users\\PPIO\\Desktop\\pictures\\a.png") // 需要替换为实际的路径
	if err != nil {
		fmt.Println("打开图片文件失败：", err.Error())
		return
	}
	defer file.Close()

	// 创建 MultipartWriter 对象
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加 image_type 字段
	writer.WriteField("image_type", "message")

	// 添加 image 文件字段
	part, err := writer.CreateFormFile("image", "test.png")
	if err != nil {
		fmt.Println("创建文件字段失败：", err.Error())
		return
	}
	io.Copy(part, file)

	// 结束写入并获取 Content-Type
	writer.Close()
	contentType := writer.FormDataContentType()

	// 创建请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("创建请求失败：", err.Error())
		return
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", "Bearer "+tToken) // 需要替换为实际的 token

	// 发送请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败：", err.Error())
		return
	}
	defer res.Body.Close()

	// 获取响应内容
	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("解析响应失败：", err.Error())
		return
	}

	imageKey := data["data"].(map[string]interface{})["image_key"].(string)

	// 构建发送的数据
	sendData := map[string]interface{}{
		"msg_type": "image",
		"content": map[string]interface{}{
			"image_key": imageKey,
		},
	}
	jsonData, err := json.Marshal(sendData)
	if err != nil {
		fmt.Println("序列化发送数据失败：", err.Error())
		return
	}

	// 发送消息
	webhookURL := "https://open.feishu.cn/open-apis/bot/v2/hook/e5a46a90-b152-4410-b2b3-522a43e202ab" // 需要替换为实际的 webhook URL
	req, err = http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("创建发送请求失败：", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	_, err = client.Do(req)
	if err != nil {
		fmt.Println("发送消息失败：", err.Error())
		return
	}

	fmt.Println("发送图片成功")
}
