package core

import (
	"fmt"
	"strings"
)

func HandleMerge(c *Context) {
	if c.SubCommand != "merge" {
		c.Next()
		return
	}

	// 禁止merge dev
	v := c.Index(2)
	if v == "dev" {
		c.Abort()
		fmt.Printf("> %s\n %s\n", c.RawCommand(), "merge dev是禁止操作，会导致分支污染。")
		return
	}
	branch := GetCurrentBranchName()

	if branch!="dev" && strings.HasPrefix(v, "dev") {
		c.Abort()
		fmt.Printf("> %s\n %s\n", c.RawCommand(), "merge dev*是禁止操作，会导致分支污染。只有dev分支允许mere dev*")
		return
	}
}
