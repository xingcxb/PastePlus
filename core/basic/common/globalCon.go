package common

const (
	// Windows Windows系统
	Windows = "windows"
	// MacOS macOS系统
	MacOS = "darwin"
	// Linux Linux系统
	Linux = "linux"
)

const (
	MainWindowName     = "MainWindowName"     // 主窗口名称
	SettingsWindowName = "SettingsWindowName" // 设置窗口名称
	UpdateWindowName   = "UpdateWindowName"   // 更新软件窗口名称
)

const (
	MainWindowContentUrl             = "#/home"              // 主窗口内容
	SettingsWindowContentUrl         = "#/settings"          // 设置窗口内容
	SettingsGeneralWindowContentUrl  = "#/settings/general"  // 设置通用窗口内容
	SettingsShortcutWindowContentUrl = "#/settings/shortcut" // 设置快捷键窗口内容
	SettingsUpdateWindowContentUrl   = "#/settings/update"   // 设置更新窗口内容
	SettingsAboutWindowContentUrl    = "#/settings/about"    // 设置关于窗口内容
	UpdateWindowUrl                  = "#/update"            // 更新窗口
)

const (
	EventsFindPasteHistoryToCore           = "findPasteHistoryToCore"          // 查找历史剪贴板数据，从页面传递到go
	EventsFindPasteHistoryToFrontend       = "findPasteHistoryToFrontend"      // 查找历史剪贴板数据，从go到页面
	EventsHandleCardClickToCore            = "handleCardClickToCore"           // 单击卡片操作事件名称，从页面传递到go
	EventsHandleCardClickToFrontend        = "handleCardClickToFrontend"       // 单击卡片操作事件名称，从go到页面
	EventsHandleCardDoubleClickToCore      = "handleCardDoubleClickToCore"     // 双击卡片操作事件名称，从页面传递到go
	EventsHandleCardDoubleClickToFrontend  = "handleCardDoubleClickToFrontend" // 双击卡片操作事件名称，从go传递到页面
	EventsHandleBootUpToCore               = "handleBootUpToCore"              // 开机启动事件名称，从页面传递到go
	EventsHandleBootUpToFrontend           = "handleBootUpToFrontend"          // 开机启动事件名称，从go传递到页面
	EventsHandleCleanAllPasteHistoryToCore = "cleanAllPasteHistoryToCore"      // 清空所有历史剪贴板数据，从页面传递到go
	EventsHandLoadPasteConfigToCore        = "loadPasteConfigToCore"           // 加载配置文件，从页面传递到go
	EventsHandLoadPasteConfigToFrontend    = "loadPasteConfigToFrontend"       // 加载配置文件，从go传递到页面
)
