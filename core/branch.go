package core

import (
	"fmt"
	"strings"
)

func HandleBranch(c *Context) {
	if c.SubCommand != "branch" {
		c.Next()
		return
	}

	// 禁止merge dev
	v := c.Index(2)
	currentBranch := GetCurrentBranchName()
	if strings.HasPrefix(currentBranch, "dev") && len(c.Args) == 3 && !strings.Contains(v, "-") {
		c.Abort()
		fmt.Printf("> %s\n %s\n", c.RawCommand(), "禁止在祖先分支为dev的场景下，新建分支。会造成分支污染")
		return
	}
}
