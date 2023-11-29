package cli

import (
	"github.com/starter-go/application"
	"github.com/starter-go/cli"

	"github.com/starter-go/cli/gen/core4cli"
	"github.com/starter-go/cli/gen/extention4cli"
	"github.com/starter-go/cli/gen/test4cli"
)

// Module 导出模块：[github.com/starter-go/cli]
func Module() application.Module {
	mb := cli.CreateCoreModule(nil)
	mb.Components(core4cli.Components)
	return mb.Create()
}

// ModuleExtention 导出模块：[github.com/starter-go/cli#extention]
func ModuleExtention() application.Module {
	parent := Module()

	mb := cli.CreateExtentionModule(nil)
	mb.Components(extention4cli.Components)
	mb.Depend(parent)
	return mb.Create()
}

// ModuleTest 导出模块：[github.com/starter-go/cli#test]
func ModuleTest() application.Module {
	parent1 := Module()
	parent2 := ModuleExtention()

	mb := cli.CreateTestModule(nil)
	mb.Components(test4cli.Components)
	mb.Depend(parent1, parent2)
	return mb.Create()
}
