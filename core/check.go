package core

import "fmt"

func Check(c *Context) {

	if c.Index(1) == "version" ||
		c.Index(1) == "--version" || c.Index(1) == "clone" || c.Index(1)=="init"{
		c.Next()
		return
	}

	if !IsCorrectGitDir() {
		c.Abort()
		fmt.Println("未在当前目录找到.git目录，请确认是合法git项目，或使用git init 初始化")
		return
	}
}
