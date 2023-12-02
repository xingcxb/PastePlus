package kit

import (
	"PastePlus/core/basic/common"
	"github.com/xingcxb/goKit/core/strKit"
	"strconv"
)

// CheckPid 检查pid，如果不一致，返回upPid
/*
 * @param upPid 上一个应用的pid
 * @param appPid 当前应用的pid
 */
func CheckPid(upPid, appPid int) int {
	common.Logger.Info(strKit.Splicing("当前激活的窗口的Pid：", strconv.Itoa(upPid), "当前应用的Pid：", strconv.Itoa(appPid)))
	if upPid == appPid {
		return 0
	} else {
		return upPid
	}
}
