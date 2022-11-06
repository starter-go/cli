package cli

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName     = "bitwormhole.com/starter/cli"
	theModuleVersion  = "v0.0.9"
	theModuleRevision = 9
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// InitModule 初始化一个 ModuleBuilder
func InitModule(mb *application.ModuleBuilder) *application.ModuleBuilder {
	if mb == nil {
		mb = &application.ModuleBuilder{}
	}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))
	mb.Dependency(starter.Module())
	return mb
}
