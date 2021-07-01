package core

func GetCurrentBranchName() string {
	rs := GetCmdOutPut("git rev-parse --abbrev-ref HEAD")
	return rs
}
