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

	newb := c.Index(3)
	if strings.HasPrefix(currentBranch, "dev") && b == "-b" && !strings.HasPrefix(newb, "dev") {
		c.Abort()
		fmt.Printf("> %s\n %s\n", c.RawCommand(), "禁止在祖先分支为dev的场景下，新建功能分支。会造成分支污染")
		return
	}
}
