package libcli

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/gen/cfgcommon"
	"bitwormhole.com/starter/cli/gen/cfgexample"
	"bitwormhole.com/starter/cli/gen/cfglib"
	"github.com/bitwormhole/starter/application"
)

// Module ... 导出模块[bitwormhole.com/starter/cli]
func Module() application.Module {
	mb := &application.ModuleBuilder{}
	mb = cli.InitModule(mb)
	mb.OnMount(cfglib.ExportConfig)
	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleCommon ... 导出模块[bitwormhole.com/starter/cli#common]
func ModuleCommon() application.Module {

	parent := Module()

	mb := &application.ModuleBuilder{}
	mb.Name(parent.GetName() + "#common")
	mb.Version(parent.GetVersion())
	mb.Revision(parent.GetRevision())

	mb.Resources(parent.GetResources())
	mb.Dependency(parent)
	mb.OnMount(cfgcommon.ExportConfig)
	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleExample ... 导出模块[bitwormhole.com/starter/cli#example]
func ModuleExample() application.Module {

	parent := ModuleCommon()

	mb := &application.ModuleBuilder{}
	mb.Name(parent.GetName() + "#example")
	mb.Version(parent.GetVersion())
	mb.Revision(parent.GetRevision())

	mb.Resources(parent.GetResources())
	mb.Dependency(parent)
	mb.OnMount(cfgexample.ExportConfig)
	return mb.Create()
}
