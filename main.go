package main

import (
	"fmt"
	"github.com/fwhezfwhez/gitx/core"
	"os"
	"os/exec"
)

var Version = "gitx version v1.1.0"

func main() {
	//  os.Args = []string{"sudo","gitx", "--version"}

	args := core.GetRealArgs(os.Args)

	ctx := core.NewContext(args)

	// fmt.Println(core.Debug(ctx))

	f := func() {
		cmd := exec.Command(core.GetGitCommandPath(), args[1:]...)

		rs, e := cmd.CombinedOutput()
		if e != nil {
			// panic(e)
		}

		if args[1] == "--version" || args[1] == "version" {
			fmt.Println(Version)
			fmt.Println(string(rs))
			return
		}

		fmt.Println(string(rs))
		return
	}

	wrapf := core.WrapFuncWithContext(f, ctx)

	// 预检查
	wrapf.Use(core.Check)

	// 切面操作merge
	wrapf.Use(core.HandleMerge)
	// 切面操作branch
	wrapf.Use(core.HandleBranch)
	// 切面操作checkout
	wrapf.Use(core.HandleCheckout)
	// 切面操作add
	core.AddRouter(wrapf)

	wrapf.Handle()
}
