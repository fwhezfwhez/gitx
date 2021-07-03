package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

// 获取当前的分支名
func GetCurrentBranchName() string {
	rs := GetCmdOutPut("git rev-parse --abbrev-ref HEAD")
	return Trim(rs)
}

// 获取当前hash
func GetCurrentHash() string {
	rs := GetCmdOutPut("git rev-parse HEAD")
	return Trim(rs)
}

// 获取git命令
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

// 获取当前状态是否处于merging中
// 返回bool, git路径
func IsOnMergingState() (bool, string) {
	dir, e := os.Getwd()
	if e != nil {
		panic(e)
	}

	gitDir := path.Join(dir, ".git")
	files, e := ioutil.ReadDir(gitDir)
	if e != nil {
		panic(e)
	}

	for _, v := range files {
		if strings.HasPrefix(v.Name(), "MERGE") {
			return true, gitDir
		}
	}
	return false, gitDir

}

// 是否为处理冲突类型的分支
func IsCftTypeBranch() bool {
	branchName := GetCurrentBranchName()

	if strings.Contains(branchName, "cft") {
		return true
	}
	return false
}

func IsCorrectGitDir() bool {
	dir, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	gitdir := path.Join(dir, ".git")
	info, e := os.Stat(gitdir)
	if os.IsNotExist(e) {
		return false
	}
	if e != nil {
		panic(e)
	}

	if info.IsDir() == true {
		return true
	}

	return false

}
