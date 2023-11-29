package core4cli
import (
    p1336d65ed "github.com/starter-go/cli"
    pb6de49a96 "github.com/starter-go/cli/implements/core"
     "github.com/starter-go/application"
)

// type pb6de49a96.ComCLI in package:github.com/starter-go/cli/implements/core
//
// id:com-b6de49a969de63e2-core-ComCLI
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:alias-1336d65edeed550b78a5d5b61e92d726-CLI
// scope:singleton
//
type pb6de49a969_core_ComCLI struct {
}

func (inst* pb6de49a969_core_ComCLI) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b6de49a969de63e2-core-ComCLI"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = "alias-1336d65edeed550b78a5d5b61e92d726-CLI"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb6de49a969_core_ComCLI) new() any {
    return &pb6de49a96.ComCLI{}
}

func (inst* pb6de49a969_core_ComCLI) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb6de49a96.ComCLI)
	nop(ie, com)

	
    com.Handlers = inst.getHandlers(ie)
    com.Filters = inst.getFilters(ie)


    return nil
}


func (inst*pb6de49a969_core_ComCLI) getHandlers(ie application.InjectionExt)[]p1336d65ed.HandlerRegistry{
    dst := make([]p1336d65ed.HandlerRegistry, 0)
    src := ie.ListComponents(".class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry")
    for _, item1 := range src {
        item2 := item1.(p1336d65ed.HandlerRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pb6de49a969_core_ComCLI) getFilters(ie application.InjectionExt)[]p1336d65ed.FilterRegistry{
    dst := make([]p1336d65ed.FilterRegistry, 0)
    src := ie.ListComponents(".class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p1336d65ed.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}


