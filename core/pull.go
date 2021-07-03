package core

import "fmt"

func PullRouter(r *WrapF) {
	// 只允许同名分支pull
	r.Use(pullSameBranchOnly)
}

func pullSameBranchOnly(c *Context) {
	// 如果不是pull类命令，不处理
	if c.SubCommand != "pull" {
		c.Next()
		return
	}

	remoteBranch := c.Index(3)

	localBranch := GetCurrentBranchName()

	if remoteBranch!= localBranch {
		fmt.Println("禁止跨分支pull,只允许同名分支发起pull")
		c.Abort()
		return
	}

}