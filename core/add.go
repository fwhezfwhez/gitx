package core

import (
	"fmt"
	"strings"
)

func AddRouter(f *WrapF) {
	// 限制在cft类型分支里，非merging状态时，add时确认提交信息，功能类驳回
	f.Use(addLimitCftMergeStateAdd)
}

// 限制在cft类型分支里，非merging状态时，add时确认提交信息，功能类驳回
func addLimitCftMergeStateAdd(c *Context) {
	// 如果不是add类命令，不处理
	if c.SubCommand != "add" {
		c.Next()
		return
	}

	// 如果不是cft类型分支，放行
	if !IsCftTypeBranch() {
		return
	}

	// merging中允许add。会被理解为处理冲突的add
	isMerging, _ := IsOnMergingState()
	if isMerging {
		return
	}

	// 确认add的内容类型
	fmt.Printf("在cft类型分支上提交内容为异常操作，请确认提交内容是为 追加冲突处理【Y】，还是功能开发【N】。\n")
	fmt.Printf("Y/N: ")

	var yn string
	fmt.Scan(&yn)

	// 对处理冲突类型操作放行
	if strings.ToUpper(yn) == "Y" || strings.ToUpper(yn) == "YES" {
		fmt.Println("允许追加处理冲突。")
		c.Next()
		return
	}

	// 对功能类型操作拦截并阻止
	if strings.ToUpper(yn) == "N" || strings.ToUpper(yn) == "NO" {
		c.Abort()
		fmt.Println("禁止在冲突处理分支中，追加功能开发，add已被撤销。")
		return
	}

	// 对非法回答
	c.Abort()
	fmt.Println("非法回答Y/N.操作已撤销!")
	return

}
