package db

import (
	"PastePlus/core/basic"
	"PastePlus/core/basic/common"
	"PastePlus/core/kit"
	"PastePlus/core/plugin/dialogKit"
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xingcxb/goKit/core/dateKit"
	"github.com/xingcxb/goKit/core/fileKit"
	"github.com/xingcxb/goKit/core/strKit"
	"go.uber.org/zap"
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
			common.Logger.Error("创建数据库文件失败", zap.Error(err))
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "创建数据库文件失败")
			return
		}
		sqlite3Db, err = sql.Open("sqlite3", filePath)
		if err != nil {
			common.Logger.Error("数据库初始化失败", zap.Error(err))
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "数据库初始化失败")
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
		if err = basicConfigInit(sqlite3Db); err != nil {
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "初始化数据库配置表失败")
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

// 初始化基础配置表
func basicConfigInit(sqlite3Db *sql.DB) error {
	// 初始化配置表
	sqlStmt := "INSERT INTO pasteConfig(key, value) VALUES(?, ?)"
	// 设置启动状态为false
	_, err := sqlite3Db.Exec(sqlStmt, common.ConfigKeyBootUp, "false")
	if err != nil {
		return err
	}
	_, err = sqlite3Db.Exec(sqlStmt, common.ConfigKeySound, "false")
	if err != nil {
		return err
	}
	_, err = sqlite3Db.Exec(sqlStmt, common.ConfigKeyHistoryCapacity, "周")
	if err != nil {
		return err
	}
	return nil
}

// UpConfig 更新配置信息
func UpConfig(config common.PasteConfig) bool {
	sqlStm := "UPDATE pasteConfig SET value = ? WHERE id = ?"
	exec, err := sqlite3Db.Exec(sqlStm, config.Value, config.Id)
	if err != nil {
		common.Logger.Error("更新配置信息失败", zap.Error(err))
		return false
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		common.Logger.Error("更新配置信息失败", zap.Error(err))
		return false
	}
	if affected > 1 {
		common.Logger.Error("更新配置信息条数错误", zap.Error(err))
		return true
	}
	if affected == 1 {
		return true
	}
	return false
}

// FindAllConfig 获取所有的配置信息数据
func FindAllConfig() ([]common.PasteConfig, error) {
	sqlStms := "SELECT * FROM pasteConfig"
	rows, err := sqlite3Db.Query(sqlStms)
	if err != nil {
		return nil, err
	}
	var pasteConfigList []common.PasteConfig
	for rows.Next() {
		var pasteConfig common.PasteConfig
		_ = rows.Scan(&pasteConfig.Id, &pasteConfig.Key, &pasteConfig.Value)
		pasteConfigList = append(pasteConfigList, pasteConfig)
	}
	return pasteConfigList, nil
}

func FindConfigByKey(key string) (common.PasteConfig, error) {
	pasteConfig := common.PasteConfig{}
	sqlStms := "SELECT * FROM pasteConfig WHERE key = ?"
	rows, err := sqlite3Db.Query(sqlStms, key)
	if err != nil {
		//dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库配置表失败")
		//fmt.Println("查询数据库配置表失败", err.Error())
		return pasteConfig, err
	}
	for rows.Next() {
		_ = rows.Scan(&pasteConfig.Id, &pasteConfig.Key, &pasteConfig.Value)
	}
	return pasteConfig, nil
}

// ResetPasteHistoryData 重置剪贴板历史数据
func ResetPasteHistoryData() (bool, error) {
	// 删除所有的数据
	sqlStm := "DELETE FROM pasteHistory"
	_, err := sqlite3Db.Exec(sqlStm)
	if err != nil {
		return false, err
	}
	// 重置id
	sqlStm = "UPDATE sqlite_sequence SET seq = 0 WHERE name = 'pasteHistory'"
	_, err = sqlite3Db.Exec(sqlStm)
	if err != nil {
		return false, err
	}
	return true, nil
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
		common.Logger.Error("查询数据库存储消息表失败", zap.Error(err))
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
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

// FindPasteListByLTDate 指定日期查询小于当前日期的数据
/*
 * @param gtDate 最小时间
 * @return 返回数据集合
 */
func FindPasteListByLTDate(gtDate string) []common.PasteHistoryGo {
	sqlStms := "SELECT * FROM pasteHistory WHERE created_at < ?"
	rows, err := sqlite3Db.Query(sqlStms, strKit.Splicing(gtDate, " 23:59:59"))
	if err != nil {
		common.Logger.Error("查询数据库存储消息表失败", zap.Error(err))
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
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
		common.Logger.Error("查询数据库存储消息表失败", zap.Error(err))
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
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
		common.Logger.Error("查询数据库存储消息表失败", zap.Error(err))
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
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
		common.Logger.Error("查询数据库存储消息表失败", zap.Error(err))
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
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
		common.Logger.Error("查询数据库存储消息表失败", zap.Error(err))
		dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "查询数据库存储消息表失败")
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
			common.Logger.Error("保存剪贴板数据失败", zap.Error(err))
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "保存剪贴板数据失败")
			return false
		}
		return true
	} else {
		// 更新创建时间
		sqlStm := "UPDATE pasteHistory SET created_at = ? WHERE id = ?"
		_, err := sqlite3Db.Exec(sqlStm, dateKit.Now(), data.Id)
		if err != nil {
			common.Logger.Error("更新剪贴板数据失败", zap.Error(err))
			dialogKit.PackageTipsDialog(dialogKit.Error, "错误", "更新剪贴板数据失败")
			return false
		}
		return true
	}
}

// DeletePasteById 通过id删除数据
/*
 * @param id 数据id
 */
func DeletePasteById(id int64) error {
	sqlStm := "DELETE FROM pasteHistory WHERE id = ?"
	_, err := sqlite3Db.Exec(sqlStm, id)
	if err != nil {
		common.Logger.Error("删除数据失败", zap.Error(err))
		return err
	}
	return nil
}
