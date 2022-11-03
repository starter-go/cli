package main

import (
	"bitwormhole.com/starter/cli/libcli"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.UseMain(libcli.ModuleExample())
	i.Run()
}
