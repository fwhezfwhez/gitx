package core

import (
	"fmt"
	"testing"
)

func TestGetRealArgs(t *testing.T) {
	rs := GetRealArgs([]string{
		"sudo", "gitx", "branch", "-v",
	})

	fmt.Println(rs)
}
