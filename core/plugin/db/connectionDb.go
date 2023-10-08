package db

import (
	"PastePlus/core/basic"
	"PastePlus/core/basic/common"
	"PastePlus/core/kit"
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/xingcxb/goKit/core/fileKit"
	"github.com/xingcxb/goKit/core/strKit"
	"os"
)

// CheckDbFile 检查数据库文件是否存在
func CheckDbFile(filePath string) bool {
	return fileKit.Exists(filePath)
}

// 用于判断是否存在sqlite3的文件，不存在就创建，存在就不创建
func InitDb() (*sql.DB, error) {
	homeDir, _ := kit.HomeDir(context.Background())
	filePath := strKit.Splicing(homeDir+string(os.PathSeparator), basic.AppConfigPath, string(os.PathSeparator), basic.AppDbPath)
	if !CheckDbFile(filePath) {
		// 当文件不存在的时候创建数据库文件，并创建表
		err := kit.CreateLazyFile(filePath)
		if err != nil {
			return nil, err
		}
		db, err := sql.Open("sqlite3", filePath)
		if err != nil {
			errDialog := application.ErrorDialog()
			errDialog.SetTitle("错误")
			errDialog.SetMessage("数据库初始化失败")
			errDialog.Show()
			return nil, err
		}
		// 创建存储数据表
		sqlStmt := "CREATE TABLE IF NOT EXISTS pasteHistory (id INTEGER PRIMARY KEY AUTOINCREMENT,from_app TEXT, content BLOB,type TEXT, created_at TEXT)"
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, err
		}
		// 创建配置表
		sqlStmt = "CREATE TABLE IF NOT EXISTS pasteConfig (id INTEGER PRIMARY KEY AUTOINCREMENT,key TEXT, value TEXT)"
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, err
		}
		return db, nil
	} else {
		// 当文件存在的时候直接打开数据库文件
		db, err := sql.Open("sqlite3", filePath)
		if err != nil {
			errDialog := application.ErrorDialog()
			errDialog.SetTitle("错误")
			errDialog.SetMessage("数据库初始化失败")
			errDialog.Show()
			return nil, err
		}
		return db, nil
	}
}

// UpConfig 更新配置信息
func UpConfig(config common.PasteConfig) bool {

	return false
}

// FindAllConfig 获取所有的配置信息数据
func FindAllConfig() []common.PasteConfig {
	db, err := InitDb()
	if err != nil {
		errDialog := application.ErrorDialog()
		errDialog.SetTitle("错误")
		errDialog.SetMessage(err.Error())
		errDialog.Show()
		return nil
	}
	sqlStms := "SELECT * FROM pasteConfig"
	rows, err := db.Query(sqlStms)
	if err != nil {
		errDialog := application.ErrorDialog()
		errDialog.SetTitle("错误")
		errDialog.SetMessage(err.Error())
		errDialog.Show()
		return nil
	}
	defer db.Close()
	var pasteConfigList []common.PasteConfig
	for rows.Next() {
		var pasteConfig common.PasteConfig
		_ = rows.Scan(&pasteConfig.Id, &pasteConfig.Key, &pasteConfig.Value)
		pasteConfigList = append(pasteConfigList, pasteConfig)
	}
	return pasteConfigList
}

// FindListByGTDate 指定日期查询大于当前日期的数据
/*
 * @param gtDate 最小时间
 * @return 返回数据集合
 */
func FindListByGTDate(gtDate string) []common.PasteHistory {
	db, err := InitDb()
	if err != nil {
		errDialog := application.ErrorDialog()
		errDialog.SetTitle("错误")
		errDialog.SetMessage(err.Error())
		errDialog.Show()
		return nil
	}
	sqlStms := "SELECT * FROM pasteHistory WHERE created_at > ?"
	rows, err := db.Query(sqlStms, strKit.Splicing(gtDate, " 00:00:00"))
	if err != nil {
		errDialog := application.ErrorDialog()
		errDialog.SetTitle("错误")
		errDialog.SetMessage(err.Error())
		errDialog.Show()
		return nil
	}
	defer db.Close()
	var pasteHistoryList []common.PasteHistory
	for rows.Next() {
		var pasteHistory common.PasteHistory
		_ = rows.Scan(&pasteHistory.Id, &pasteHistory.FromApp, &pasteHistory.Content, &pasteHistory.Type, &pasteHistory.CreatedAt)
		pasteHistoryList = append(pasteHistoryList, pasteHistory)
	}
	return pasteHistoryList
}

// FindListByDate 指定日期查询数据
/*
 * @param beginDate 开始时间
 * @param endDate 结束时间
 * @return 返回数据集合
 */
func FindListByDate(beginDate, endDate string) []common.PasteHistory {
	beginTime := strKit.Splicing(beginDate, " 00:00:00")
	endTime := strKit.Splicing(endDate, " 23:59:59")
	db, err := InitDb()
	if err != nil {
		errDialog := application.ErrorDialog()
		errDialog.SetTitle("错误")
		errDialog.SetMessage(err.Error())
		errDialog.Show()
		return nil
	}
	sqlStms := "SELECT * FROM pasteHistory WHERE created_at BETWEEN ? AND ? ORDER BY created_at DESC"
	rows, err := db.Query(sqlStms, beginTime, endTime)
	if err != nil {
		errDialog := application.ErrorDialog()
		errDialog.SetTitle("错误")
		errDialog.SetMessage(err.Error())
		errDialog.Show()
		return nil
	}
	defer db.Close()
	var pasteHistoryList []common.PasteHistory
	for rows.Next() {
		var pasteHistory common.PasteHistory
		_ = rows.Scan(&pasteHistory.Id, &pasteHistory.FromApp, &pasteHistory.Content, &pasteHistory.Type, &pasteHistory.CreatedAt)
		pasteHistoryList = append(pasteHistoryList, pasteHistory)
	}
	return pasteHistoryList
}
