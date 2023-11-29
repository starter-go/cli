package main

import (
	"os"

	"github.com/starter-go/cli/modules/cli"
	"github.com/starter-go/units"
)

func main() {

	m := cli.ModuleExtention()

	r := units.NewRunner()
	r.Dependencies(m)
	r.Run(os.Args)
}
