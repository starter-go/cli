package cli

import (
	"github.com/starter-go/application"
	"github.com/starter-go/cli"
)

// Module 导出模块：[github.com/starter-go/cli]
func Module() application.Module {
	mb := cli.CreateModule(nil)
	mb.Components(nil) // todo
	return mb.Create()
}

// ModuleCommon 导出模块：[github.com/starter-go/cli#common]
func ModuleCommon() application.Module {
	parent := Module()
	mb := cli.CreateModule(nil).Name(parent.Name() + "#common")
	mb.Components(nil) // todo
	mb.Depend(parent)
	return mb.Create()
}

// ModuleExample 导出模块：[github.com/starter-go/cli#example]
func ModuleExample() application.Module {
	parent1 := Module()
	parent2 := ModuleCommon()
	mb := cli.CreateModule(nil).Name(parent1.Name() + "#example")
	mb.Components(nil) // todo
	mb.Depend(parent1, parent2)
	return mb.Create()
}
