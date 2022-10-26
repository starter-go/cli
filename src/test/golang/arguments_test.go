package gunit

import (
	"fmt"
	"testing"

	"bitwormhole.com/starter/cli"
)

func TestArguments(t *testing.T) {

	rows := []string{
		"",
		"abc",
		"abc def",
		"   git config --global user.email 'foo@example.com'    ",
		"   git commit     -m='demo: hello, world '     ",
	}

	argsTmp := cli.LineToArguments(rows[4])
	t.Log(argsTmp)

	for _, line1 := range rows {
		args := cli.LineToArguments(line1)
		line2 := cli.ArgumentsToLine(args)
		fmt.Println("=========================================")
		fmt.Println("line_1: ", line1)
		fmt.Println("line_2: ", line2)
		fmt.Println("args:")
		for i, a := range args {
			fmt.Println("  args[", i, "] = `", a, "`")
		}
	}
}
