package basic

import _ "embed"

//go:embed version.txt
var AppVersion string // AppVersion 应用版本

const (
	// AppName 应用名称
	AppName = "PastePlus"
	// AppDescription 应用描述
	AppDescription = "PastePlus 是一款跨平台的剪贴板管理工具，支持 Windows、MacOS、Linux。"
	// AppConfigPath 应用配置文件夹名称
	AppConfigPath = ".PastePlus"
	// AppDbPath 应用数据库文件名称
	AppDbPath = "pastePlus.db"
	// AppAuthorSite 应用程序作者网站
	AppAuthorSite = "https://xingcxb.com?from=PastePlus"
	// AppHelpCenter 应用程序帮助中心
	AppHelpCenter = "https://paste.xingcxb.com"
	// AppGithubIssues 应用程序Github的issues地址
	AppGithubIssues = "https://github.com/xingcxb/PastePlus/issues"
)

const (
	ConfigKeyBootUp          = "bootUp"          // 是否开机启动标识
	ConfigKeyPasteText       = "pasteText"       // 是否粘贴为纯文本标识
	ConfigKeySound           = "sound"           // 是否开启声音标识
	ConfigKeyMenuIcon        = "menuIcon"        // 是否开启菜单栏图标标识
	ConfigKeyHistoryCapacity = "historyCapacity" // 历史记录容量标识
)
