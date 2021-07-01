package main

import (
	"fmt"
	"gitx/core"
	"os"
	"os/exec"
)

func main() {
	// os.Args = []string{"sudo","gitx", "merge", "dev"}

	args := core.GetRealArgs(os.Args)

	ctx := core.NewContext(args)

	// fmt.Println(core.Debug(ctx))

	f := func() {
		cmd := exec.Command("git", args[1:]...)

		rs, e := cmd.CombinedOutput()
		if e != nil {
			// panic(e)
		}

		fmt.Println(string(rs))
		return
	}

	wrapf := core.WrapFuncWithContext(f, ctx)

	// 切面操作merge
	wrapf.Use(core.HandleMerge)
	// 切面操作branch
	wrapf.Use(core.HandleBranch)
	// 切面操作checkout
	wrapf.Use(core.HandleCheckout)

	wrapf.Handle()
}
