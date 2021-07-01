package core

import (
	"fmt"
	"strings"
)

func HandleCheckout(c *Context) {
	if c.SubCommand != "checkout" {
		c.Next()
		return
	}

	// 禁止merge dev
	b := c.Index(2)
	currentBranch := GetCurrentBranchName()
	if strings.HasPrefix(currentBranch, "dev") && b == "-b" {
		c.Abort()
		fmt.Printf("> %s\n %s\n", c.RawCommand(), "禁止在祖先分支为dev的场景下，新建分支。会造成分支污染")
		return
	}
}
