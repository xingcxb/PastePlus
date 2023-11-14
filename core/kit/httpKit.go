package kit

import (
	"errors"
	"github.com/xingcxb/goKit/core/strKit"
	"io"
	"net/http"
	"os"
	"strings"
)

// HttpDownloadWithProgress 带进度条的下载
/*
 * @param urlString 网址
 * @param savePath 保存路径 末尾是否携带/都可以
 * @param fileName 文件名，如果不存在则自动获取
 * @param isCover 是否覆盖 true 覆盖 false 不覆盖(当文件存在的时候返回该文件已存在)
 * @param fn func(rate int) 进度回调函数
 * @return string 文件路径,error
 */
func HttpDownloadWithProgress(urlString, savePath, fileName string, isCover bool, fn func(rate int)) (string, error) {
	if savePath == "" {
		return "", errors.New("保存路径为空")
	}
	// 发起网络请求
	// 必须要优先请求的原因是使用的开发人员可能没有指定文件名称，需要从url中获取
	res, err := http.Get(urlString)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if fileName == "" {
		//定义文件名字
		path := strings.Split(urlString, "/")
		fileName = path[len(path)-1]
	}
	// 检查保存路径是否以/结尾
	if !strings.HasSuffix(savePath, "/") {
		savePath = strKit.Splicing(savePath, "/")
	}

	filePath := strKit.Splicing(savePath, fileName)
	// 判断文件是否存在，默认不存在
	checkFile := false
	// 判断文件是否存在
	if _, err := os.Stat(filePath); err == nil {
		checkFile = true
	}
	if isCover && checkFile {
		// 允许覆盖，删除文件
		err = os.Remove(filePath)
		if err != nil {
			return "", err
		}
	} else if isCover {
		// 不允许覆盖，返回错误
		if checkFile {
			return "", errors.New("文件已存在")
		}
	}
	//创建文件
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	// defer延迟调用 关闭文件，释放资源
	defer out.Close()
	// 获取文件大小
	fileSize := res.ContentLength
	buf := make([]byte, 1024)
	counter := &WriteCounter{
		Expected: uint64(fileSize),
		Callback: fn,
	}
	// 使用 io.TeeReader 在写入文件的同时更新计数器
	source := io.TeeReader(res.Body, counter)
	// 使用 io.CopyBuffer 写入文件
	_, err = io.CopyBuffer(out, source, buf)
	if err != nil {
		return "", err
	}
	fn(100) // 调用进度回调函数
	return filePath, nil
}

type WriteCounter struct {
	Total    uint64
	Expected uint64
	Callback func(int)
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc *WriteCounter) PrintProgress() {
	percent := int(float64(wc.Total) / float64(wc.Expected) * 100)
	wc.Callback(percent)
}
