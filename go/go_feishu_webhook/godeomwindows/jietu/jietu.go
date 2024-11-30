package jietu

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

// 截图函数
func Jietufunlunc() { // Chrome 驱动路径和选项
	chromeOptions := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		//chromedp.Headless, //无头模式
		chromedp.WindowSize(1920, 3450),
		chromedp.DisableGPU,
		chromedp.NoSandbox,
	}

	// 目标 URL
	xlcloudcywURL := "目标rl"

	// 创建 Chrome 实例
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), chromeOptions...)
	defer cancel()

	// 创建 Chromedp 上下文
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// 启动 Chrome 实例
	err := chromedp.Run(ctx)
	if err != nil {
		fmt.Printf("无法启动 Chrome 实例：%s\n", err.Error())
		return
	}

	// 打开目标页面
	err = chromedp.Run(ctx, chromedp.Navigate(xlcloudcywURL))
	if err != nil {
		fmt.Printf("无法打开网页：%s\n", err.Error())
		return
	}

	//time.Sleep(5 * time.Second)
	//fmt.Println("已经等待了5秒了")
	// 等待页面加载完成
	err = chromedp.Run(ctx, chromedp.WaitReady("body"))

	if err != nil {
		fmt.Printf("页面加载超时：%s\n", err.Error())
		return
	}

	err = chromedp.Run(ctx,
		chromedp.WaitVisible(`input[name="user"]`),                    // 等待用户名输入框可见
		chromedp.SendKeys(`input[name="user"]`, "business@localhost"), // 替换成你的用户名
		chromedp.SendKeys(`input[name="password"]`, "business@123"),   // 替换成你的密码
		chromedp.Click(`button[class="css-14g7ilz-button"]`),          // 替换成登录按钮的选择器
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("已经等待5秒了")

	// 调整视窗大小
	viewport := chromedp.EmulateViewport(1920, 1080)
	err = chromedp.Run(ctx, viewport)
	if err != nil {
		fmt.Printf("无法调整视口大小：%s\n", err.Error())
		return
	}

	// 截取整个页面的截图
	var buf []byte
	err = chromedp.Run(ctx, chromedp.CaptureScreenshot(&buf))
	if err != nil {
		fmt.Printf("无法截图：%s\n", err.Error())
		return
	}

	// 保存截图到文件
	file, err := os.Create("C:\\Users\\PPIO\\Desktop\\pictures\\a.png")
	if err != nil {
		fmt.Printf("无法创建截图文件：%s\n", err.Error())
		return
	}
	defer file.Close()

	_, err = file.Write(buf)
	if err != nil {
		fmt.Printf("无法保存截图：%s\n", err.Error())
		return
	}

	fmt.Println("操作完成")

}

// 指定截图区域
//x := 70
//y := 65
//width := 1920
//height := 1020
//
//// 滚动页面使截图区域可见
//err = chromedp.Run(ctx, chromedp.ScrollIntoView(fmt.Sprintf(`document.elementFromPoint(%d, %d)`, x, y)))
//if err != nil {
//	fmt.Errorf("无法滚动页面：%s\n", err.Error())
//}
//
//// 截取整个可视区域的截图
//var buf []byte
//err = chromedp.Run(ctx, chromedp.CaptureScreenshot(&buf))
//if err != nil {
//	log.Fatal(err)
//}
//
//// 将整个截图进行裁剪，获取指定区域的截图
//decodedImg, err := png.Decode(bytes.NewReader(buf))
//if err != nil {
//	log.Fatal(err)
//}
//croppedImg := imaging.Crop(decodedImg, image.Rect(x, y, x+width, y+height))

// 保存截图到文件
//file, err := os.Create("C:\\Users\\PPIO\\Desktop\\pictures\\a.png")
//if err != nil {
//	fmt.Printf("无法创建截图文件：%s\n", err.Error())
//	return
//}
//defer file.Close()
//
//err = png.Encode(file, croppedImg)
//if err != nil {
//	fmt.Printf("无法保存截图：%s\n", err.Error())
//	return
//}
//
//fmt.Println("操作完成")

//}

// 下面在单独是个函数 然后在调用这样

//func Jietufunlunc2() {
//	//调整视窗大小
//	//err = chromedp.EmulateViewport(1920, 1080).Do(ctx)
//	//if err != nil {
//	//	fmt.Printf("无法调整视口大小：%s\n", err.Error())
//	//	return
//	//}
//	//ctx, cancel := chromedp.NewContext(context.Background())
//	//defer cancel()
//
//
//}

//// 截图成功的函数
//func jietusuccess() {
//
//}

//暂时废弃的代码

//err= chromedp.Run(ctx,chromedp.)
//
//
//chromedp.ActionFunc(func(ctx context.Context) error { // 执行自定义的操作函数
//	// get layout metrics
//	_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx) // // 获取网页布局的度量信息
//	if err != nil {
//		return err
//	}
//
//	width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))
//
//
//	// 截取整个页面的截图
//	var buf []byte
//	*res, err = page.CaptureScreenshot(). // // 调用ChromeDP的CaptureScreenshot()方法，返回一个截图结果对象
//		WithClip(&page.Viewport{
//			X:      contentSize.X,      // // 裁剪区域的左上角横坐标
//			Y:      contentSize.Y,      //// 裁剪区域的左上角纵坐标
//			Width:  contentSize.Width,  /// 裁剪区域的宽度
//			Height: contentSize.Height, //// 裁剪区域的高度
//			Scale:  1,                  // // 缩放比例
//		}).Do(ctx) //// 执行截图操作
//
//	)
//	if err != nil {
//		fmt.Printf("无法截图：%s\n", err.Error())
//		return
//	}
