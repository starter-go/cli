package gunit

import (
	"fmt"
	"testing"

	"github.com/starter-go/cli/arguments"
)

func TestArguments(t *testing.T) {

	rows := []string{
		"",
		"abc",
		"abc def",
		"   git config --global user.email 'foo@example.com'    ",
		"   git commit     -m='demo: hello, world '     ",
	}

	argsTmp := arguments.Parse(rows[4])
	t.Log(argsTmp)

	for _, line1 := range rows {
		args := arguments.Parse(line1)
		line2 := arguments.Assemble(args)
		fmt.Println("=========================================")
		fmt.Println("line_1: ", line1)
		fmt.Println("line_2: ", line2)
		fmt.Println("args:")
		for i, a := range args {
			fmt.Println("  args[", i, "] = `", a, "`")
		}
	}
}

func TestArgumentsTemplate(t *testing.T) {

	args := []string{"git", "clone", "-l", "-s", "hello", "-b", "name1"}

	atb := arguments.NewTemplateBuilder()
	atb.AcceptOption("--template", 1) // --template=<template-directory>
	atb.AcceptOption("-l", 0)
	atb.AcceptOption("-s", 0)
	atb.AcceptOption("--no-hardlinks", 0)
	atb.AcceptOption("-q", 0)
	atb.AcceptOption("-n", 0)
	atb.AcceptOption("--bare", 0)
	atb.AcceptOption("-o", 1)          // -o <name>
	atb.AcceptOption("-b", 1)          // -b <name>
	atb.AcceptOption("-u", 1)          // -u <upload-pack>
	atb.AcceptOption("--reference", 1) // --reference <repository>
	at := atb.Create()

	aa := at.Parse(args)
	vlist := make(map[int]any, 0)

	vlist[0] = aa.GetItem(0).String()
	vlist[1] = aa.GetOption("--abc").Exists()
	vlist[2] = aa.GetOption("--xyz").Value(2).Exists()
	vlist[3] = aa.GetOption("-b").Value(0).String()

	for k, v := range vlist {
		fmt.Println(k, v)
	}

}
