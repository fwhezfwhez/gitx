package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args

	cmd := exec.Command("git", args[1:]...)

	rs, e := cmd.Output()
	if e != nil {
		panic(e)
	}

	fmt.Println(string(rs))
}
