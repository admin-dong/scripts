package main

import (
	"fmt"
	"godeom/jietu"
	token2 "godeom/token"
	"godeom/webhook"
)

// 截图保存路径

func main() {
	jietu.Jietufunlunc()

	token, err := token2.T_token()
	if err != nil {
		fmt.Println("获取token失败", err)
	}
	webhook.SendImage(token)
	if err != nil {
		fmt.Println()
	}

}
