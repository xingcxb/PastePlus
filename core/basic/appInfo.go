package basic

const (
	// AppName 应用名称
	AppName = "Paste+"
	// AppVersion 应用版本
	AppVersion = "0.0.1"
	// AppConfigPath 应用配置文件夹名称
	AppConfigPath = ".PastePlus"
	// AppDbPath 应用数据库文件名称
	AppDbPath = "pastePlus.db"
)

const (
	ConfigKeyBootUp          = "bootUp"          // 是否开机启动标识
	ConfigKeyPasteText       = "pasteText"       // 是否粘贴为纯文本标识
	ConfigKeySound           = "sound"           // 是否开启声音标识
	ConfigKeyMenuIcon        = "menuIcon"        // 是否开启菜单栏图标标识
	ConfigKeyHistoryCapacity = "historyCapacity" // 历史记录容量标识
)
