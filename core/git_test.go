package core

import (
	"fmt"
	"testing"
)

func TestGetCurrentBranchName(t *testing.T) {
	rs := GetCurrentBranchName()

	fmt.Println(rs)
}
