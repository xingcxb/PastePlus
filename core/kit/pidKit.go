package kit

// CheckPid 检查pid，如果不一致，返回upPid
/*
 * @param upPid 上一个应用的pid
 * @param appPid 当前应用的pid
 */
func CheckPid(upPid, appPid int) int {
	if upPid == appPid {
		return 0
	} else {
		return upPid
	}
}
