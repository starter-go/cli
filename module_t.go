package cli

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/cli"
	theModuleVersion  = "v1.0.1"
	theModuleRevision = 11

	theCoreModuleResPath      = "src/core/resources"
	theTestModuleResPath      = "src/test/resources"
	theExtentionModuleResPath = "src/extention/resources"
)

//go:embed "src/core/resources"
var theCoreModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

//go:embed "src/extention/resources"
var theExtentionModuleResFS embed.FS

// CreateCoreModule 初始化一个 ModuleBuilder
func CreateCoreModule(mb *application.ModuleBuilder) *application.ModuleBuilder {
	if mb == nil {
		mb = &application.ModuleBuilder{}
	}
	mb.Name(theModuleName + "#core")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theCoreModuleResFS, theCoreModuleResPath)
	mb.Depend(starter.Module())
	return mb
}

// CreateExtentionModule 初始化一个 ModuleBuilder
func CreateExtentionModule(mb *application.ModuleBuilder) *application.ModuleBuilder {
	if mb == nil {
		mb = &application.ModuleBuilder{}
	}
	mb.Name(theModuleName + "#extention")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theExtentionModuleResFS, theExtentionModuleResPath)
	mb.Depend(starter.Module())
	return mb
}

// CreateTestModule 初始化一个 ModuleBuilder
func CreateTestModule(mb *application.ModuleBuilder) *application.ModuleBuilder {
	if mb == nil {
		mb = &application.ModuleBuilder{}
	}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	mb.Depend(starter.Module())
	return mb
}
