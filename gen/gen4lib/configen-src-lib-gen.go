package gen4lib
import (
    p1336d65ed "github.com/starter-go/cli"
    p510644c9c "github.com/starter-go/cli/config/lib"
     "github.com/starter-go/application"
)

// type p510644c9c.ComCLI in package:github.com/starter-go/cli/config/lib
//
// id:com-510644c9c919a104-lib-ComCLI
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:alias-1336d65edeed550b78a5d5b61e92d726-CLI
// scope:singleton
//
type p510644c9c9_lib_ComCLI struct {
}

func (inst* p510644c9c9_lib_ComCLI) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-510644c9c919a104-lib-ComCLI"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = "alias-1336d65edeed550b78a5d5b61e92d726-CLI"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p510644c9c9_lib_ComCLI) new() any {
    return &p510644c9c.ComCLI{}
}

func (inst* p510644c9c9_lib_ComCLI) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p510644c9c.ComCLI)
	nop(ie, com)

	
    com.Handlers = inst.getHandlers(ie)
    com.Filters = inst.getFilters(ie)


    return nil
}


func (inst*p510644c9c9_lib_ComCLI) getHandlers(ie application.InjectionExt)[]p1336d65ed.HandlerRegistry{
    dst := make([]p1336d65ed.HandlerRegistry, 0)
    src := ie.ListComponents(".class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry")
    for _, item1 := range src {
        item2 := item1.(p1336d65ed.HandlerRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p510644c9c9_lib_ComCLI) getFilters(ie application.InjectionExt)[]p1336d65ed.FilterRegistry{
    dst := make([]p1336d65ed.FilterRegistry, 0)
    src := ie.ListComponents(".class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p1336d65ed.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p510644c9c.Filters in package:github.com/starter-go/cli/config/lib
//
// id:com-510644c9c919a104-lib-Filters
// class:class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry
// alias:
// scope:singleton
//
type p510644c9c9_lib_Filters struct {
}

func (inst* p510644c9c9_lib_Filters) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-510644c9c919a104-lib-Filters"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p510644c9c9_lib_Filters) new() any {
    return &p510644c9c.Filters{}
}

func (inst* p510644c9c9_lib_Filters) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p510644c9c.Filters)
	nop(ie, com)

	


    return nil
}


