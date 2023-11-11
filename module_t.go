package cli

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/cli"
	theModuleVersion  = "v0.0.9+"
	theModuleRevision = 9
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// CreateModule 初始化一个 ModuleBuilder
func CreateModule(mb *application.ModuleBuilder) *application.ModuleBuilder {
	if mb == nil {
		mb = &application.ModuleBuilder{}
	}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Depend(starter.Module())
	return mb
}
