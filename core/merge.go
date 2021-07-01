package core

import "fmt"

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
}
