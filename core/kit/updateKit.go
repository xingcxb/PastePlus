// Package kit Description: 更新工具包
package kit

import (
	"PastePlus/core/basic"
	"PastePlus/core/basic/common"
	"PastePlus/core/plugin/dialogKit"
	"context"
	"github.com/ncruces/zenity"
	"github.com/tidwall/gjson"
	"github.com/xingcxb/goKit/core/httpKit"
	"github.com/xingcxb/goKit/core/strKit"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var (
	ownerGithub = "xingcxb"   // GitHub 仓库的所有者
	repoGithub  = "PastePlus" // GitHub 仓库的名称
)

// UpdateAppVersion 更新应用版本
func UpdateAppVersion() {
	// 检查最新版本
	releaseInfo := CheckLatestVersion()
	if releaseInfo.Version == basic.AppVersion {
		dialogKit.PackageTipsDialog(dialogKit.Info, "提示", "🎉 当前已经是最新版本")
		return
	}
	// 创建一个提示框
	err := zenity.Question(strKit.Splicing("检测到新版本：", releaseInfo.Version, "\n", releaseInfo.Changelog),
		zenity.Title("提示"),
		zenity.Icon(zenity.InfoIcon),
		zenity.OKLabel("更新"),
		zenity.CancelLabel("取消"))
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "唤起更新提示框失败")
		return
	}
	// 创建一个进度条
	progress, err := zenity.Progress(zenity.Title("下载更新"), zenity.MaxValue(100))
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "唤起下载进度条失败")
		return
	}
	// 执行更新应用
	updateApp(releaseInfo, progress)
}

// updateApp 更新应用
func updateApp(releaseInfo common.ReleaseInfo, progress zenity.ProgressDialog) {
	// 获取系统主目录
	homeDir, err := HomeDir(context.Background())
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "获取系统主目录失败")
		return
	}
	// 下载地址
	url := ""
	// 文件名称
	fileName := "PastePlus"
	// 获取系统类型
	sysType := runtime.GOOS
	switch sysType {
	case common.MacOS:
		// Mac
		url = releaseInfo.DownloadURLs[common.MacOS]
		fileName = strKit.Splicing(fileName, ".dmg")
	case common.Linux:
		// Linux
		url = releaseInfo.DownloadURLs[common.Linux]
		fileName = strKit.Splicing(fileName, ".AppImage")
	case common.Windows:
		// Windows
		url = releaseInfo.DownloadURLs[common.Windows]
		fileName = strKit.Splicing(fileName, ".exe")
	default:
		// 其他
		dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "暂不支持该系统")
		return
	}
	// 发起网络请求
	filePath, err := HttpDownloadWithProgress(url, strKit.Splicing(homeDir, "/Downloads/"), fileName, true, func(rate int) {
		progress.Text(strKit.Splicing("下载进度：", strconv.Itoa(rate), "%"))
		progress.Value(rate)
	})
	progress.Text("下载完成")
	progress.Value(100)
	progress.Close()
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "下载失败")
		return
	}
	// 更新应用核心
	updateAppCore(filePath, sysType)
}

// updateAppCore 更新应用核心
func updateAppCore(filePath, sysType string) {
	switch sysType {
	case common.MacOS:
		// Mac
		cmd := exec.Command("open", filePath)
		err := cmd.Run()
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "更新失败")
			return
		}
	case common.Windows:
		// Windows
		dialogKit.PackageTipsDialog(dialogKit.Info, "🎉 下载完成", "请在下载目录安装")
	case common.Linux:
		// Linux
		dialogKit.PackageTipsDialog(dialogKit.Info, "🎉 下载完成", "请在下载目录安装")
	}
}

// CheckLatestVersion 检查最新版本
func CheckLatestVersion() common.ReleaseInfo {
	return CheckVersionInGithub()
}

// CheckVersionInGithub 检查GitHub
func CheckVersionInGithub() common.ReleaseInfo {
	releaseInfo := common.ReleaseInfo{}
	url := strKit.Splicing("https://api.github.com/repos/", ownerGithub, "/", repoGithub, "/releases")
	// 发起网络请求
	responseStr, err := httpKit.HttpGet(url)
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Warning, "错误", "检查最新版本失败")
		return releaseInfo // 发生错误，直接返回
	}
	if responseStr == "[]" {
		dialogKit.PackageTipsDialog(dialogKit.Info, "提示", "🎉 当前已经是最新版本")
		return releaseInfo
	}
	lastResult := gjson.Parse(responseStr).Array()[0]
	// 版本号
	releaseInfo.Version = lastResult.Get("tag_name").String()
	// 发布时间
	releaseInfo.ReleaseTime = lastResult.Get("published_at").String()
	// 更新日志
	releaseInfo.Changelog = lastResult.Get("body").String()
	// 资源文件
	assets := lastResult.Get("assets").Array()
	downloadMap := make(map[string]string, 0)
	for _, asset := range assets {
		// 获取下载链接
		downloadURL := asset.Get("browser_download_url").String()
		// 获取应用名称
		downloadName := asset.Get("name").String()
		if strings.Contains(downloadName, "exe") {
			// Windows平台
			downloadMap[common.Windows] = downloadURL
		} else if strings.Contains(downloadName, "dmg") {
			// macOS平台
			downloadMap[common.MacOS] = downloadURL
		} else {
			// Linux平台
			downloadMap[common.Linux] = downloadURL
		}
	}
	// 下载地址
	releaseInfo.DownloadURLs = downloadMap
	return releaseInfo
}
