package gen4example
import (
    p1336d65ed "github.com/starter-go/cli"
    pd3796ea0f "github.com/starter-go/cli/config/example"
     "github.com/starter-go/application"
)

// type pd3796ea0f.TestPoint in package:github.com/starter-go/cli/config/example
//
// id:com-d3796ea0fc4e66e2-example-TestPoint
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type pd3796ea0fc_example_TestPoint struct {
}

func (inst* pd3796ea0fc_example_TestPoint) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d3796ea0fc4e66e2-example-TestPoint"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd3796ea0fc_example_TestPoint) new() any {
    return &pd3796ea0f.TestPoint{}
}

func (inst* pd3796ea0fc_example_TestPoint) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd3796ea0f.TestPoint)
	nop(ie, com)

	
    com.CLI = inst.getCLI(ie)


    return nil
}


func (inst*pd3796ea0fc_example_TestPoint) getCLI(ie application.InjectionExt)p1336d65ed.CLI{
    return ie.GetComponent("#alias-1336d65edeed550b78a5d5b61e92d726-CLI").(p1336d65ed.CLI)
}


