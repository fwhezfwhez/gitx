package core

import (
	"fmt"
	"testing"
)

func TestGetCurrentBranchName(t *testing.T) {
	rs := GetCurrentBranchName()

	fmt.Println(rs)
}

func TestGetGitCommandPath(t *testing.T) {
	path := GetGitCommandPath()

	fmt.Println(path)
}

func TestGetCurrentHash(t *testing.T) {
	rs := GetCurrentHash()

	fmt.Println(rs)
}
