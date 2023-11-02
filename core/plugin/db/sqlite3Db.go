package db

import (
	"PastePlus/core/basic"
	"PastePlus/core/basic/common"
	"PastePlus/core/kit"
	"PastePlus/core/plugin/dialogKit"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xingcxb/goKit/core/dateKit"
	"github.com/xingcxb/goKit/core/fileKit"
	"github.com/xingcxb/goKit/core/strKit"
	"os"
	"time"
)

// sqlite3Db 数据库连接
var (
	sqlite3Db *sql.DB // sqlite3数据库连接
	// Sqlite3Status 启动状态，默认为false。true：启动，false：未启动
	Sqlite3Status = false
)

// CheckDbFile 检查数据库文件是否存在
func CheckDbFile(filePath string) bool {
	return fileKit.Exists(filePath)
}

// InitDb 初始化sqlite3 用于判断是否存在sqlite3的文件，不存在就创建，存在就不创建
func InitDb() {
	// 延时启动数据库，让主窗口先启动
	time.Sleep(300 * time.Millisecond)
	homeDir, _ := kit.HomeDir(context.Background())
	filePath := strKit.Splicing(homeDir+string(os.PathSeparator), basic.AppConfigPath, string(os.PathSeparator), basic.AppDbPath)
	if !CheckDbFile(filePath) {
		// 当文件不存在的时候创建数据库文件，并创建表
		err := kit.CreateLazyFile(filePath)
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "创建数据库文件失败")
			fmt.Println("创建数据库文件失败", err)
			return
		}
		sqlite3Db, err = sql.Open("sqlite3", filePath)
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "数据库初始化失败")
			fmt.Println("数据库初始化失败", err)
			return
		}
		// 创建存储数据表
		sqlStmt := "CREATE TABLE IF NOT EXISTS pasteHistory (id INTEGER PRIMARY KEY AUTOINCREMENT,from_app TEXT, content BLOB,type TEXT, created_at TEXT)"
		_, err = sqlite3Db.Exec(sqlStmt)
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "创建数据库存储消息表失败")
			return
		}
		// 创建配置表
		sqlStmt = "CREATE TABLE IF NOT EXISTS pasteConfig (id INTEGER PRIMARY KEY AUTOINCREMENT,key TEXT, value TEXT)"
		_, err = sqlite3Db.Exec(sqlStmt)
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "创建数据库配置表失败")
			return
		}
	} else {
		// 当文件存在的时候直接打开数据库文件
		db, err := sql.Open("sqlite3", filePath)
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "数据库初始化失败")
			return
		}
		sqlite3Db = db
	}
	// 设置启动状态为true
	Sqlite3Status = true
}

// UpConfig 更新配置信息
func UpConfig(config common.PasteConfig) bool {

	return false
}

// FindAllConfig 获取所有的配置信息数据
func FindAllConfig() []common.PasteConfig {
	sqlStms := "SELECT * FROM pasteConfig"
	rows, err := sqlite3Db.Query(sqlStms)
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库配置表失败")
		fmt.Println("查询数据库配置表失败", err.Error())
		return nil
	}
	var pasteConfigList []common.PasteConfig
	for rows.Next() {
		var pasteConfig common.PasteConfig
		_ = rows.Scan(&pasteConfig.Id, &pasteConfig.Key, &pasteConfig.Value)
		pasteConfigList = append(pasteConfigList, pasteConfig)
	}
	return pasteConfigList
}

// FindPasteListByGTDate 指定日期查询大于当前日期的数据
/*
 * @param gtDate 最小时间
 * @return 返回数据集合
 */
func FindPasteListByGTDate(gtDate string) []common.PasteHistoryGo {
	sqlStms := "SELECT * FROM pasteHistory WHERE created_at > ?"
	rows, err := sqlite3Db.Query(sqlStms, strKit.Splicing(gtDate, " 00:00:00"))
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
		fmt.Println("查询数据库存储消息表失败", err.Error())
		return nil
	}
	var pasteHistoryList []common.PasteHistoryGo
	for rows.Next() {
		var pasteHistory common.PasteHistoryGo
		_ = rows.Scan(&pasteHistory.Id, &pasteHistory.FromApp, &pasteHistory.Content, &pasteHistory.Type, &pasteHistory.CreatedAt)
		pasteHistoryList = append(pasteHistoryList, pasteHistory)
	}
	return pasteHistoryList
}

// FindPasteListByDate 指定日期查询数据
/*
 * @param beginDate 开始时间
 * @param endDate 结束时间
 * @return 返回数据集合
 */
func FindPasteListByDate(beginDate, endDate string) []common.PasteHistoryGo {
	beginTime := strKit.Splicing(beginDate, " 00:00:00")
	endTime := strKit.Splicing(endDate, " 23:59:59")
	sqlStms := "SELECT * FROM pasteHistory WHERE created_at BETWEEN ? AND ? ORDER BY created_at DESC"
	rows, err := sqlite3Db.Query(sqlStms, beginTime, endTime)
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
		fmt.Println("查询数据库存储消息表失败", err.Error())
		return nil
	}
	var pasteHistoryList []common.PasteHistoryGo
	for rows.Next() {
		var pasteHistory common.PasteHistoryGo
		_ = rows.Scan(&pasteHistory.Id, &pasteHistory.FromApp, &pasteHistory.Content, &pasteHistory.Type, &pasteHistory.CreatedAt)
		pasteHistoryList = append(pasteHistoryList, pasteHistory)
	}
	return pasteHistoryList
}

// FindPasteById 通过id获取粘贴板历史数据
/*
 * @param id 粘贴板id
 * @return 查询到的数据
 */
func FindPasteById(id int) common.PasteHistoryGo {
	pasteHistory := common.PasteHistoryGo{}
	sqlStms := "SELECT * FROM pasteHistory WHERE id = ?"
	rows, err := sqlite3Db.Query(sqlStms, id)
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
		fmt.Println("查询数据库存储消息表失败", err.Error())
		return pasteHistory
	}
	for rows.Next() {
		_ = rows.Scan(&pasteHistory.Id, &pasteHistory.FromApp, &pasteHistory.Content, &pasteHistory.Type, &pasteHistory.CreatedAt)
	}
	return pasteHistory
}

// FindPasteDataByContent 根据内容查询数据
/*
 * @param content 内容
 * @return 返回数据
 */
func FindPasteDataByContent(content []byte) common.PasteHistoryGo {
	sqlStm := "SELECT * FROM pasteHistory WHERE content = ? ORDER BY created_at DESC LIMIT 1"
	rows, err := sqlite3Db.Query(sqlStm, content)
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
		fmt.Println("查询数据库存储消息表失败", err.Error())
		return common.PasteHistoryGo{}
	}
	var pasteHistoryList []common.PasteHistoryGo
	for rows.Next() {
		var pasteHistory common.PasteHistoryGo
		_ = rows.Scan(&pasteHistory.Id, &pasteHistory.FromApp, &pasteHistory.Content, &pasteHistory.Type, &pasteHistory.CreatedAt)
		pasteHistoryList = append(pasteHistoryList, pasteHistory)
	}
	if len(pasteHistoryList) == 0 {
		return common.PasteHistoryGo{}
	}
	return pasteHistoryList[0]
}

// FindPasteDateByAll 获取所有的剪贴板数据
func FindPasteDateByAll() []common.PasteHistoryGo {
	sqlStm := "SELECT * FROM pasteHistory ORDER BY created_at DESC"
	rows, err := sqlite3Db.Query(sqlStm)
	if err != nil {
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
		fmt.Println("查询数据库存储消息表失败", err.Error())
		return nil
	}
	var pasteHistoryList []common.PasteHistoryGo
	for rows.Next() {
		var pasteHistory common.PasteHistoryGo
		_ = rows.Scan(&pasteHistory.Id, &pasteHistory.FromApp, &pasteHistory.Content, &pasteHistory.Type, &pasteHistory.CreatedAt)
		pasteHistoryList = append(pasteHistoryList, pasteHistory)
	}
	return pasteHistoryList
}

// SaveOrUpdatePaste 保存或更新剪贴板数据
/*
 * @param data 剪贴板数据
 * @return 返回是否成功
 */
func SaveOrUpdatePaste(data common.PasteHistoryGo) bool {
	if data.Id == 0 {
		sqlStm := "INSERT INTO pasteHistory(from_app, content, type, created_at) VALUES(?, ?, ?, ?)"
		_, err := sqlite3Db.Exec(sqlStm, data.FromApp, data.Content, data.Type, data.CreatedAt)
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "保存剪贴板数据失败")
			fmt.Println("保存剪贴板数据失败", err.Error())
			return false
		}
		return true
	} else {
		// 更新创建时间
		sqlStm := "UPDATE pasteHistory SET created_at = ? WHERE id = ?"
		_, err := sqlite3Db.Exec(sqlStm, dateKit.Now(), data.Id)
		if err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "更新剪贴板数据失败")
			fmt.Println("更新剪贴板数据失败", err.Error())
			return false
		}
		return true
	}
}
