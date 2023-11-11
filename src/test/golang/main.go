package main

import (
	"os"

	"github.com/starter-go/cli/modules/cli"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(cli.ModuleExample())
	i.Run()
}
