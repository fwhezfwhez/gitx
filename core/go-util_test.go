package core

import (
	"fmt"
	"testing"
)

func TestGetCmdOutPut(t *testing.T) {
	lb := GetCmdOutPut("git rev-parse --abbrev-ref HEAD")

	fmt.Println(lb)
}


func TestGetCmdOutPut2(t *testing.T) {
	lb := GetCmdOutPut("git rev-parse --abbrev-ref HEAD")

	fmt.Println(lb)
}