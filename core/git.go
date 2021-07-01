package core

import (
	"fmt"
	"runtime"
)

// 获取当前的分支名
func GetCurrentBranchName() string {
	rs := GetCmdOutPut("git rev-parse --abbrev-ref HEAD")
	return rs
}

// 获取当前hash
func GetCurrentHash() string {
	rs := GetCmdOutPut("git rev-parse HEAD")
	return rs
}

func GetGitCommandPath() (string) {
	switch runtime.GOOS {
	case "windows":
		rs := GetCmdOutPut("where git")
		if rs == "" {
			panic(fmt.Errorf("git命令未正确配置,请前往https://git-scm.com/downloads选择正确的版本下载并安装"))
		}

		return Trim(rs)
	case "darwin", "linux":
		rs := GetCmdOutPut("which git")
		if rs == "" {
			panic(fmt.Errorf("git命令未正确配置,请前往https://git-scm.com/downloads选择正确的版本下载并安装"))
		}
		return Trim(rs)
	default:
		fmt.Println("gitx只支持windows，darwin，linux三种操作系统,已自动降级为git")
		return "git"
	}
}
