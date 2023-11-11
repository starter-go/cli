package cli

import (
	"github.com/starter-go/application"
	"github.com/starter-go/cli"
	"github.com/starter-go/cli/gen/gen4common"
	"github.com/starter-go/cli/gen/gen4example"
	"github.com/starter-go/cli/gen/gen4lib"
)

// Module 导出模块：[github.com/starter-go/cli]
func Module() application.Module {
	mb := cli.CreateModule(nil)
	mb.Components(gen4lib.ComForCLI)
	return mb.Create()
}

// ModuleCommon 导出模块：[github.com/starter-go/cli#common]
func ModuleCommon() application.Module {
	parent := Module()
	mb := cli.CreateModule(nil).Name(parent.Name() + "#common")
	mb.Components(gen4common.ComForCLICommon)
	mb.Depend(parent)
	return mb.Create()
}

// ModuleExample 导出模块：[github.com/starter-go/cli#example]
func ModuleExample() application.Module {
	parent1 := Module()
	parent2 := ModuleCommon()
	mb := cli.CreateModule(nil).Name(parent1.Name() + "#example")
	mb.Components(gen4example.ComForCLIExample)
	mb.Depend(parent1, parent2)
	return mb.Create()
}
