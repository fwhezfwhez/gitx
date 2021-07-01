package core

// 获取当前的分支名
func GetCurrentBranchName() string {
	rs := GetCmdOutPut("git rev-parse --abbrev-ref HEAD")
	return rs
}

func GetCurrentHash() string {
	rs := GetCmdOutPut("git rev-parse HEAD")
	return rs
}
