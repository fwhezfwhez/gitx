package core

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// 获取gitx开始的命令链
func GetRealArgs(args [] string) []string {
	// 定向到gitx命令
	var rs = make([]string, 0, 10)

	var start int
	for i, v := range args {
		if v == "gitx" {
			start = i
		}
	}

	rs = append(rs, args[start:]...)

	return rs
}

func Debug(i interface{}) string {
	rs, _ := json.MarshalIndent(i, "  ", "  ")
	return string(rs)
}

func GetValueOfIndex(arr []string, index int) string {
	if len(arr) == 0 {
		return ""
	}

	if index < len(arr) {
		return arr[index]
	}
	return ""
}

// 获取一段原生命令的返回结果
func GetCmdOutPut(command string) (string) {
	arr := strings.Split(command, " ")

	cmd := exec.Command(arr[0], arr[1:]...)
	rs, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Println(e.Error())
	}
	return string(rs)
}
